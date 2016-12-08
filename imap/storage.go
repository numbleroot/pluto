package imap

import (
	"fmt"
	"log"
	"net"
	"os"

	"crypto/tls"
	"path/filepath"

	"github.com/numbleroot/maildir"
	"github.com/numbleroot/pluto/comm"
	"github.com/numbleroot/pluto/config"
	"github.com/numbleroot/pluto/crdt"
	"github.com/numbleroot/pluto/crypto"
)

// Structs

// Storage struct bundles information needed in
// operation of a storage node.
type Storage struct {
	Socket           net.Listener
	MailboxStructure map[string]map[string]*crdt.ORSet
	ApplyCRDTUpdChan chan string
	DoneCRDTUpdChan  chan bool
	Config           *config.Config
}

// Functions

// InitStorage listens for TLS connections on a TCP socket
// opened up on supplied IP address. It returns those
// information bundeled in above Storage struct.
func InitStorage(config *config.Config) (*Storage, error) {

	var err error

	// Initialize and set fields.
	storage := &Storage{
		MailboxStructure: make(map[string]map[string]*crdt.ORSet),
		ApplyCRDTUpdChan: make(chan string),
		DoneCRDTUpdChan:  make(chan bool),
		Config:           config,
	}

	// Find all files below this node's CRDT root layer.
	userFolders, err := filepath.Glob(filepath.Join(config.Storage.CRDTLayerRoot, "*"))
	if err != nil {
		return nil, fmt.Errorf("[imap.InitStorage] Globbing for CRDT folders of users failed with: %s\n", err.Error())
	}

	for _, folder := range userFolders {

		// Retrieve information about accessed file.
		folderInfo, err := os.Stat(folder)
		if err != nil {
			return nil, fmt.Errorf("[imap.InitStorage] Error during stat'ing possible user CRDT folder: %s\n", err.Error())
		}

		// Only consider folders for building up CRDT map.
		if folderInfo.IsDir() {

			// Extract last part of path, the user name.
			userName := filepath.Base(folder)

			// Read in mailbox structure CRDT from file.
			userMainCRDT, err := crdt.InitORSetFromFile(filepath.Join(folder, "mailbox-structure.log"))
			if err != nil {
				return nil, fmt.Errorf("[imap.InitStorage] Reading CRDT failed: %s\n", err.Error())
			}

			// Store main CRDT in designated map for user name.
			storage.MailboxStructure[userName] = make(map[string]*crdt.ORSet)
			storage.MailboxStructure[userName]["Structure"] = userMainCRDT

			// Retrieve all mailboxes the user possesses
			// according to main CRDT.
			userMailboxes := userMainCRDT.GetAllValues()

			for _, userMailbox := range userMailboxes {

				// Read in each mailbox CRDT from file.
				userMailboxCRDT, err := crdt.InitORSetFromFile(filepath.Join(folder, fmt.Sprintf("%s.log", userMailbox)))
				if err != nil {
					return nil, fmt.Errorf("[imap.InitStorage] Reading CRDT failed: %s\n", err.Error())
				}

				// Store each read-in CRDT in map under
				storage.MailboxStructure[userName][userMailbox] = userMailboxCRDT
			}
		}
	}

	// Load internal TLS config.
	internalTLSConfig, err := crypto.NewInternalTLSConfig(config.Storage.TLS.CertLoc, config.Storage.TLS.KeyLoc, config.RootCertLoc)
	if err != nil {
		return nil, err
	}

	// Start to listen for incoming internal connections on defined IP and sync port.
	storage.Socket, err = tls.Listen("tcp", fmt.Sprintf("%s:%s", config.Storage.IP, config.Storage.SyncPort), internalTLSConfig)
	if err != nil {
		return nil, fmt.Errorf("[imap.InitStorage] Listening for internal TLS connections failed with: %s\n", err.Error())
	}

	// Initialize receiving goroutine for sync operations.
	// TODO: Storage has to iterate over all worker nodes it is serving
	//       as CRDT backend for and create a 'CRDT-subnet' for each.
	_, _, err = comm.InitReceiver("storage", filepath.Join(config.Storage.CRDTLayerRoot, "receiving.log"), storage.Socket, storage.ApplyCRDTUpdChan, storage.DoneCRDTUpdChan, []string{"worker-1"})
	if err != nil {
		return nil, err
	}

	log.Printf("[imap.InitStorage] Listening for incoming sync requests on %s.\n", storage.Socket.Addr())

	return storage, nil
}

// ApplyCRDTUpd receives strings representing CRDT
// update operations from receiver and executes them.
func (storage *Storage) ApplyCRDTUpd() error {

	for {

		// Receive update message from receiver
		// via channel.
		updMsg := <-storage.ApplyCRDTUpdChan

		// Parse operation that payload specifies.
		op, opPayload, err := comm.ParseOp(updMsg)
		if err != nil {
			return fmt.Errorf("[imap.ApplyCRDTUpd] Error while parsing operation from sync message: %s\n", err.Error())
		}

		// Depending on received operation,
		// parse remaining payload further.
		switch op {

		case "create":

			// Parse received payload message into create message struct.
			createUpd, err := comm.ParseCreate(opPayload)
			if err != nil {
				return fmt.Errorf("[imap.ApplyCRDTUpd] Error while parsing CREATE update from sync message: %s\n", err.Error())
			}

			// Save user's mailbox structure CRDT to more
			// conveniently use it hereafter.
			userMainCRDT := storage.MailboxStructure[createUpd.User]["Structure"]

			// Create a new Maildir on stable storage.
			posMaildir := maildir.Dir(filepath.Join(storage.Config.Storage.MaildirRoot, createUpd.User, createUpd.AddMailbox.Value))

			err = posMaildir.Create()
			if err != nil {
				return fmt.Errorf("[imap.ApplyCRDTUpd] Maildir for new mailbox could not be created: %s\n", err.Error())
			}

			// Construct path to new CRDT file.
			posMailboxCRDTPath := filepath.Join(storage.Config.Storage.CRDTLayerRoot, createUpd.User, fmt.Sprintf("%s.log", createUpd.AddMailbox.Value))

			// Initialize new ORSet for new mailbox.
			posMailboxCRDT, err := crdt.InitORSetWithFile(posMailboxCRDTPath)
			if err != nil {

				// Perform clean up.

				log.Printf("[imap.ApplyCRDTUpd] Fail: %s\n", err.Error())
				log.Printf("[imap.ApplyCRDTUpd] Removing just created Maildir completely...\n")

				// Attempt to remove Maildir.
				err = posMaildir.Remove()
				if err != nil {
					log.Printf("[imap.ApplyCRDTUpd] ... failed to remove Maildir: %s\n", err.Error())
					log.Printf("[imap.ApplyCRDTUpd] Exiting.\n")
				} else {
					log.Printf("[imap.ApplyCRDTUpd] ... done. Exiting.\n")
				}

				// Exit worker.
				os.Exit(1)
			}

			// Write new mailbox' file to stable storage.
			err = posMailboxCRDT.WriteORSetToFile()
			if err != nil {

				// Perform clean up.

				log.Printf("[imap.ApplyCRDTUpd] Fail: %s\n", err.Error())
				log.Printf("[imap.ApplyCRDTUpd] Removing just created Maildir completely...\n")

				// Attempt to remove Maildir.
				err = posMaildir.Remove()
				if err != nil {
					log.Printf("[imap.ApplyCRDTUpd] ... failed to remove Maildir: %s\n", err.Error())
					log.Printf("[imap.ApplyCRDTUpd] Exiting.\n")
				} else {
					log.Printf("[imap.ApplyCRDTUpd] ... done. Exiting.\n")
				}

				// Exit worker.
				os.Exit(1)
			}

			// If succeeded, add a new folder in user's main CRDT.
			userMainCRDT.AddEffect(createUpd.AddMailbox.Value, createUpd.AddMailbox.Tag, true)

			// Write updated CRDT to stable storage.
			err = userMainCRDT.WriteORSetToFile()
			if err != nil {

				// Perform clean up.

				log.Printf("[imap.ApplyCRDTUpd] Fail: %s\n", err.Error())
				log.Printf("[imap.ApplyCRDTUpd] Deleting just added mailbox from main structure CRDT...\n")

				// Not optimal solution but at least keeps local version
				// consistent: remove folder from user's main CRDT.
				rSet := make(map[string]string)
				rSet[createUpd.AddMailbox.Tag] = createUpd.AddMailbox.Value

				userMainCRDT.RemoveEffect(rSet, true)

				log.Printf("[imap.ApplyCRDTUpd] ... done.\n")
				log.Printf("[imap.ApplyCRDTUpd] Removing just created Maildir completely...\n")

				// Attempt to remove Maildir.
				err = posMaildir.Remove()
				if err != nil {
					log.Printf("[imap.ApplyCRDTUpd] ... failed to remove Maildir: %s\n", err.Error())
				} else {
					log.Printf("[imap.ApplyCRDTUpd] ... done.\n")
				}

				log.Printf("[imap.ApplyCRDTUpd] Removing just created CRDT file...\n")

				// Attempt to remove just created CRDT file.
				err = os.Remove(posMailboxCRDTPath)
				if err != nil {
					log.Printf("[imap.ApplyCRDTUpd] ... failed to remove Maildir: %s\n", err.Error())
					log.Printf("[imap.ApplyCRDTUpd] Exiting.\n")
				} else {
					log.Printf("[imap.ApplyCRDTUpd] ... done. Exiting.\n")
				}

				// Exit worker.
				os.Exit(1)
			}

		case "delete":
			deleteUpd, err := comm.ParseDelete(opPayload)
			if err != nil {
				return fmt.Errorf("[imap.ApplyCRDTUpd] Error while parsing DELETE update from sync message: %s\n", err.Error())
			}

			log.Printf("APPLY HERE: DELETE %#v\n", deleteUpd.RmvMailbox)

		case "rename":
			renameUpd, err := comm.ParseRename(opPayload)
			if err != nil {
				return fmt.Errorf("[imap.ApplyCRDTUpd] Error while parsing RENAME update from sync message: %s\n", err.Error())
			}

			log.Printf("APPLY HERE: RENAME %#v\n", renameUpd.RmvMailbox)

		case "append":
			appendUpd, err := comm.ParseAppend(opPayload)
			if err != nil {
				return fmt.Errorf("[imap.ApplyCRDTUpd] Error while parsing APPEND update from sync message: %s\n", err.Error())
			}

			log.Printf("APPLY HERE: APPEND %#v\n", appendUpd.AddMail)

		case "expunge":
			expungeUpd, err := comm.ParseExpunge(opPayload)
			if err != nil {
				return fmt.Errorf("[imap.ApplyCRDTUpd] Error while parsing EXPUNGE update from sync message: %s\n", err.Error())
			}

			log.Printf("APPLY HERE: EXPUNGE %#v\n", expungeUpd.RmvMails)

		case "store":
			storeUpd, err := comm.ParseStore(opPayload)
			if err != nil {
				return fmt.Errorf("[imap.ApplyCRDTUpd] Error while parsing STORE update from sync message: %s\n", err.Error())
			}

			log.Printf("APPLY HERE: STORE %#v\n", storeUpd.AddMail)

		case "copy":
			copyUpd, err := comm.ParseCopy(opPayload)
			if err != nil {
				return fmt.Errorf("[imap.ApplyCRDTUpd] Error while parsing COPY update from sync message: %s\n", err.Error())
			}

			log.Printf("APPLY HERE: COPY %#v\n", copyUpd.AddMails)

		}

		// Signal receiver that update was performed.
		storage.DoneCRDTUpdChan <- true
	}
}
