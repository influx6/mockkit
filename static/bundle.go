package static


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

type fileData struct{
  path string
  root string
  data []byte
}

var (
  assets = map[string][]string{
    
      ".tml": []string{  // all .tml assets.
        
          "impl.tml",
        
          "mock.tml",
        
      },
    
  }

  assetFiles = map[string]fileData{
    
      
        "impl.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xbd\x8a\xe3\x30\x10\xee\xf5\x14\x5f\x71\x85\x0d\x87\xd2\x1f\x5c\x91\xe6\x20\x70\xbb\xc5\xbe\x81\x90\xc7\xb1\x40\x1e\x1b\x49\x26\x04\x31\xef\xbe\x28\xb6\x93\x38\x3f\xc5\x96\x1e\x7d\xf3\xfd\x8d\x73\xfe\x75\xe0\x44\xa1\x35\x96\x3e\x4d\x4f\xf8\xf3\x17\x7a\x3b\x11\x51\x6a\xb7\x43\xce\xdb\xb9\xc8\xa1\x1f\x3d\x1a\x6a\x1d\x53\x84\x81\x1d\xd8\x06\x4a\x84\x98\xc2\x64\x13\x4e\x9d\xb3\x1d\x5c\x3f\x7a\xea\x89\x53\x44\xea\x08\x3d\xa5\x6e\x68\x22\xda\x21\x94\xef\xd7\xc4\x70\xeb\xb7\xc6\xde\xfb\xeb\xd2\xc9\x79\x8f\xd1\xb0\xb3\xbf\x11\x07\x98\xa6\x01\x93\xa5\x18\x4d\x38\xcf\x3b\x6c\x3c\xfc\x70\x74\x56\xab\x74\x1e\xe9\x9d\xe9\xd9\x61\x56\x00\x90\x33\x82\xe1\x23\x41\x7f\x2c\x32\x22\xcb\x83\x9e\x57\xfe\x4d\x6c\xd1\x4e\x6c\xab\x9c\xf5\x3e\x1c\xa7\x12\xe7\xbf\x8b\x09\x29\x4c\x24\x52\xa3\x3c\x7c\x51\x9a\x02\xdf\x8f\x57\x7a\xe2\xa6\x70\x8a\x52\x2f\xb5\xe6\x0a\xd6\xe4\xdb\xba\x1e\xcf\x23\xa2\xaf\xd8\xaa\x5e\x8a\xb9\x94\xf9\x8c\x2c\x49\xb5\x2a\xbe\x51\x15\xda\x37\x90\xfa\xa6\xfe\x36\x5f\xce\xae\x05\x0f\x89\xfa\x31\x9d\x51\x3d\x66\xad\xe7\xd5\xe7\x06\x72\x26\x6e\x44\xd6\xa2\x5d\x8b\x99\x61\x81\x16\xd5\x58\xf0\x4b\xe3\xc5\xa5\xde\xd4\x7e\xef\xe8\x0e\xbd\x56\x4b\x3e\xd2\xed\x5a\x4f\xa4\xe5\x67\xfe\x21\x67\xb8\x70\xbc\x64\x5b\x35\x4b\x22\x25\xea\x7a\xd8\xef\x00\x00\x00\xff\xff\x81\x93\x02\x7f\x43\x03\x00\x00"),
          path: "impl.tml",
          root: "impl.tml",
        },
      
        "mock.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x57\xdf\x8b\xe3\x36\x10\x7e\x5f\xd8\xff\x61\xee\xe8\x83\x5d\x52\x67\xfb\x7a\x25\x0f\x47\x9b\x85\x42\x6f\x29\xdd\x83\x3e\x2c\x4b\xd1\xda\xe3\x58\xac\x2d\x19\x49\x4e\x36\x15\xfe\xdf\xcb\x48\x72\xfc\x23\x76\xc8\xb5\x14\xda\xf3\x4b\xf0\x68\x34\x33\xfa\xf4\xcd\x37\xce\xed\x8d\xb5\xdf\xfc\x2c\x0c\xaa\x9c\xa5\xf8\xc0\x2a\x84\x0f\x1b\x48\xc6\x96\xb6\xbd\xbd\x21\x47\x50\x4c\xec\x10\x92\x4f\x68\x0a\x99\x69\x67\x5f\xaf\xc1\xbf\xfe\xc8\xca\xf2\x5e\x2a\x6b\x13\xda\xd3\xb6\x90\x61\xce\x05\x6a\x60\x60\x8e\x35\xc2\xa1\xe0\x69\x01\x85\x2c\x33\x0d\x15\x1a\xf6\x5d\x86\x86\xf1\x52\x03\x7b\x91\x8d\x01\x53\x20\xec\xf8\x9e\x8b\x1d\xa4\xac\x24\xb3\xd6\x32\xe5\xcc\x60\xe6\x92\x1c\xb8\x29\x9c\xd3\xb4\xde\xb6\x4d\x4e\x39\xa3\x98\x42\x17\x32\x4b\x6e\x6f\x5c\xd2\x85\xd2\xb4\x51\x4d\x6a\xec\xed\x0d\x00\xc0\xef\x05\x0a\x30\xbc\xc2\xe4\x33\xaf\xd0\xdb\x1e\x0d\x53\x66\x6a\xdc\x8a\x6c\x68\xf2\xc6\xf5\x1a\x7e\x0a\x07\x91\x39\xd4\x4c\xf0\x14\x78\x0e\xba\x49\x0b\x90\x69\xda\x28\x9d\x78\xc7\x5f\x69\xe9\xd1\xb0\xf4\x15\x9e\x9e\x5f\x8e\x06\x07\xe6\xad\x52\x52\x01\xef\x8e\x65\xdb\x41\xf4\x8f\x6a\xd7\x54\x28\x0c\xec\x59\xd9\x60\x17\xad\xbf\x8c\x8f\x6a\xe7\x6f\xc2\x9b\x53\x56\x73\xc3\x4a\xfe\x27\x42\x77\x5a\x6b\x93\xed\xdb\xe7\x63\x8d\xbd\x17\x8a\xac\x1d\x26\xf9\x0d\x4d\xa3\xc4\x52\x0a\xbf\xfa\xe5\x59\x20\xa4\x71\x99\x88\x28\xcc\xa4\x45\x77\x1e\x0d\x2a\x84\x35\xaa\xc1\x75\xce\x4a\x8d\x84\x5c\xad\xe4\x9e\x67\xa8\x40\x9a\x02\x55\x34\x7f\x83\x31\xb0\x53\x18\x17\xda\x57\x0e\x15\x65\x00\x53\x70\x0d\xf8\xc6\xb5\x21\x3a\xb1\x33\x00\xf3\x46\xa4\x10\x55\x4b\xf4\x88\x27\x85\x46\xae\x92\x45\xe7\x17\x29\x4b\xb0\x17\xaf\x85\x1e\x9e\xc3\x3b\x85\x79\x89\xa9\xe1\x52\x24\x2e\xc5\xb6\x44\xca\x10\x55\x98\xcc\x61\xba\xf2\x18\x2c\xac\x11\x6c\xb1\xed\x13\xd0\xe3\x21\x05\x07\x66\xbf\x32\xba\x90\x53\x51\xc1\x97\xc2\xb8\x2b\x1a\x2c\x87\xfb\xb2\x36\x99\xb4\xda\x27\x99\xbe\xce\x37\x36\xaf\x6a\x7f\x18\xb2\xfb\xfe\x3a\xf5\xac\x0b\x56\x05\xd1\xc8\xa5\x0a\x7d\x3c\x0d\x0e\x4c\x43\xce\x91\x04\xc2\xc7\x64\x65\x29\x0f\x1a\x8e\xb2\xe9\x58\xd1\xa7\x61\x84\x22\xf5\x9c\x0b\x6e\x0a\xd4\x08\x74\xaf\xde\x6c\xe4\x69\x47\x5e\xe2\x1b\x7f\x29\x11\x0c\x3a\x3e\x74\xd2\xb0\x70\xb8\x91\x34\xcc\x4b\x9e\x5f\x09\xf7\xd0\xb3\x42\xc3\xd3\xf3\x3c\x47\xce\xe0\xbf\xac\xa7\xbd\x4e\x0d\x40\xbd\x5a\xfa\x06\x00\x9f\x79\x77\xd4\xa7\xb8\xf0\xed\xb9\x03\x21\x10\xf7\xf9\x23\x6b\x93\xae\x0b\x7e\xe1\xda\x38\xae\x10\xe5\x17\xba\x72\x05\x48\x42\x16\x77\xdd\xb0\x67\xca\x89\xf9\x62\xf3\x78\xa6\x91\xab\x77\x4b\x9c\x16\x6f\xbc\xca\x3e\xc8\x43\x14\x8f\x56\xbd\x2a\x6f\x86\xce\xdd\xfe\x85\xc6\x0b\x9e\xb3\x82\xb5\x81\xe5\x0b\xea\xeb\xcf\x65\x23\x32\xd7\xe3\xde\x46\xe8\xfe\xb1\x82\x5a\x6a\x4d\xac\xa2\x13\xd1\xc4\xf4\xc9\x09\xd7\x64\x8e\x1a\x76\xa4\x03\xc3\xcd\xc9\x44\x6b\x7c\xc5\xd3\xbe\xf6\x55\x6c\x42\xb3\x0e\x57\x96\x55\x7a\xf8\x5c\xc6\x61\x54\xcf\x9c\xcf\x59\xca\x31\x4e\x93\x2c\x83\x91\x36\x89\xdd\xaf\x2c\x6f\xf4\x23\x72\x6e\xa3\x5b\x99\xd5\xbb\x33\xa9\x6b\x27\xbc\xda\x7a\xf4\x26\xb4\xe2\x79\x00\x76\x80\x76\x88\xe8\xf7\xad\x40\xf0\x72\x12\x71\xe2\xe0\x18\xaf\x93\x07\x3c\x44\xef\x85\xf4\x03\x88\xa6\x8e\x42\x5d\x4b\x41\xa2\x44\x19\xde\xc7\xe7\x12\x3b\x2b\xb0\x8f\x82\xd3\x04\xfb\x5a\x25\x36\x1c\xef\xdf\x14\xd9\xf0\x7a\x4f\x42\x47\xd5\x2e\x8b\x98\xb5\xa1\x67\x86\xe6\xff\x8d\x56\x7b\x28\xaf\x51\x6b\x6b\x79\x0e\x42\x1a\xac\x6a\x73\x84\x68\x7a\xe8\xd8\x6f\x3d\x87\x22\x7c\x26\x7e\x89\x98\x93\x63\x86\x39\x2a\x0f\x7d\x3c\xd6\x3d\x54\xca\x69\x25\xa6\x72\x8f\x2a\x8a\x7f\x70\x96\x77\x1b\xea\x32\x98\x28\x9e\x51\x2c\x75\xff\x45\x2a\xf6\x8a\x91\xff\x62\x5e\xc1\xf7\x77\x77\x77\xf1\x9c\xe3\xc6\xff\x3e\x7d\x50\x8d\x70\x6d\xee\xc4\x22\x72\xc6\xf0\xa9\xf4\x7c\x8d\x5c\xa1\xba\x42\x9a\x5c\xd4\x91\xe6\xf4\x2f\x17\x05\xc7\xe1\xb0\x34\x23\x60\x03\xac\xae\x51\x64\xd1\xa2\xcb\x2a\x84\x0f\xe1\x5a\x8a\xfb\x1f\x9e\x9f\x8e\x78\x9e\x75\x81\x5e\xe4\xae\x89\x63\xa3\x2f\xe3\xd1\x71\xef\x27\x3d\x3b\xd8\x72\x6a\x4e\x2c\xf5\x68\x2a\x9d\xe8\x3b\x70\x26\xee\xfc\x8d\xc8\xc3\x3f\x46\x97\x67\xeb\x3f\xc3\x26\x8c\x91\xb9\xca\xa7\xa3\xe2\xaf\x00\x00\x00\xff\xff\x0c\x78\xf3\xaa\xa4\x0f\x00\x00"),
          path: "mock.tml",
          root: "mock.tml",
        },
      
    
  }
)

//==============================================================================

// FilesFor returns all files that use the provided extension, returning a
// empty/nil slice if none is found.
func FilesFor(ext string) []string {
  return assets[ext]
}

// MustFindFile calls FindFile to retrieve file reader with path else panics.
func MustFindFile(path string, doGzip bool) (io.Reader, int64) {
  reader, size, err := FindFile(path, doGzip)
  if err != nil {
    panic(err)
  }

  return reader, size
}

// FindDecompressedGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindDecompressedGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, true)
}

// MustFindDecompressedGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindDecompressedGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindDecompressedGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, false)
}

// MustFindGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFile returns a io.Reader by seeking the giving file path if it exists.
func FindFile(path string, doGzip bool) (io.Reader, int64, error){
	reader, size, err := FindFileReader(path)
	if err != nil {
		return nil, size, err
	}

	if !doGzip {
		return reader, size, nil
	}

  gzr, err := gzip.NewReader(reader)
	return gzr, size, err
}

// MustFindFileReader returns bytes.Reader for path else panics.
func MustFindFileReader(path string) (*bytes.Reader, int64){
	reader, size, err := FindFileReader(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFileReader returns a io.Reader by seeking the giving file path if it exists.
func FindFileReader(path string) (*bytes.Reader, int64, error){
  item, ok := assetFiles[path]
  if !ok {
    return nil,0, fmt.Errorf("File %q not found in file system", path)
  }

  return bytes.NewReader(item.data), int64(len(item.data)), nil
}

// MustReadFile calls ReadFile to retrieve file content with path else panics.
func MustReadFile(path string, doGzip bool) string {
  body, err := ReadFile(path, doGzip)
  if err != nil {
    panic(err)
  }

  return body
}

// ReadFile attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFile(path string, doGzip bool) (string, error){
  body, err := ReadFileByte(path, doGzip)
  return string(body), err
}

// MustReadFileByte calls ReadFile to retrieve file content with path else panics.
func MustReadFileByte(path string, doGzip bool) []byte {
  body, err := ReadFileByte(path, doGzip)
  if err != nil {
    panic(err)
  }

  return body
}

// ReadFileByte attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFileByte(path string, doGzip bool) ([]byte, error){
  reader, _, err := FindFile(path, doGzip)
  if err != nil {
    return nil, err
  }

  if closer, ok := reader.(io.Closer); ok {
    defer closer.Close()
  }

  var bu bytes.Buffer

  _, err = io.Copy(&bu, reader);
  if err != nil && err != io.EOF {
   return nil, fmt.Errorf("File %q failed to be read: %+q", path, err)
  }

  return bu.Bytes(), nil
}
