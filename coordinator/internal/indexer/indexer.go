package indexer

import (
	"bufio"
	"coordinator/internal/utils"
	"fmt"
	"io"
	"os"
	"sync"
)

type Chunk struct {
	StartPosition int64
}

type Indexer struct {
	mu sync.Mutex
	ChunksByStartPosition map[int64]*Chunk
	Idexed []*Chunk
	FileSize int64
	File *os.File
	Steps uint32
}

const INDEX_STEP = int64(1024*1024*8)

func NewIndexer(file *os.File) *Indexer {
	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println("Error al leer tamaño del archivo:", err)
	}
	fileSize := fileStat.Size()
	return &Indexer{
		ChunksByStartPosition: make(map[int64]*Chunk),
		Idexed: []*Chunk{},
		FileSize: fileSize,
		Steps: calcSteps(fileSize),
		File: file,
	}
}

func calcSteps(fileSize int64) uint32 {
	return uint32((fileSize + INDEX_STEP - 1) / INDEX_STEP)
}
 
func (i *Indexer) calcStartPositionByIndex(index uint32) int64 {
	prestart := int64(0)
	if index == 0 {
		return prestart
	}

	prestart = int64(index) * INDEX_STEP

	if prestart >= i.FileSize {
		return i.FileSize
	}

	counter := utils.ScanByteCounter{}
	
	scanner := bufio.NewScanner(io.NewSectionReader(i.File, int64(prestart), INDEX_STEP))
	scanner.Split(counter.Wrap(bufio.ScanWords))
	scanner.Scan()
	
	return counter.BytesRead + int64(prestart) + 1
}

// func (i *Indexer) IndexChunk(index int) *Chunk {
// 	start := i.calcStartPositionByIndex(index)

// 	chunk, exists := i.ChunksByStartPosition[start]

// 	if !exists {
// 		chunk = &Chunk{
// 			StartPosition: start,
// 		}
// 		i.ChunksByStartPosition[start] = chunk
// 	}

// 	i.Idexed = append(i.Idexed, chunk)
// 	return chunk
// }

func (i *Indexer) IndexChunk(index uint32) *Chunk {
	start := i.calcStartPositionByIndex(index)

	chunk := &Chunk{
		StartPosition: start,
	}

	i.mu.Lock()
	defer i.mu.Unlock()
	
	i.ChunksByStartPosition[start] = chunk

	i.Idexed = append(i.Idexed, chunk)
	return chunk
}