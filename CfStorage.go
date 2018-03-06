package mcdf

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

type DirectoryEntry struct {
	dirRepository []DirectoryEntry
	stgType       uint8
	storageClsid  string
	creationDate  []byte
	modifyDate    []byte
}

func NewDirectoryEntry(name string, stageType uint8, directoryEntries []DirectoryEntry) *DirectoryEntry {

	var de DirectoryEntry
	de.dirRepository = directoryEntries
	de.stgType = stageType
	if stageType == 2 {
		de.storageClsid = "00000000000000000000000000000000"
		de.creationDate = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		de.modifyDate = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	} else if stageType == 2 {
		val := uint64(time.Now().Nanosecond())
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, val)
		de.creationDate = buf
	}
	return &de

}

func TryNew(streamName string, stageType uint8, directoryEntries []DirectoryEntry) *DirectoryEntry {
	var de DirectoryEntry
	return &de
}

type CfStorage struct {
	compoundFile   *CompoundFile
	directoryEntry *DirectoryEntry
}

func NewCfStorage(compoundFile *CompoundFile, directoryEntry *DirectoryEntry) *CfStorage {
	cf := &CfStorage{}
	cf.compoundFile = compoundFile
	cf.directoryEntry = directoryEntry
	return cf
}

func newUUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		panic("error making uid")
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func (*CfStorage) AddStream(streamName string) *CfStream {
	return nil
}