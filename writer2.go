package main

import "fmt"

type channelWriter struct {
   Channel chan byte
}

func NewChannelWriter() *channelWriter {
   return &channelWriter{
      Channel: make(chan byte, 1024),
   }
}

func (c *channelWriter) Write(p []byte) (int, error) {
   if len(p) == 0 {
      return 0, nil
   }
   go func() {
      defer close(c.Channel)
      for _, b := range p {
         c.Channel <- b
      }
   }()
   return len(p), nil
}

func main() {
   cw := NewChannelWriter()
   file, err := os.Open("./writer2.go")
   if err != nil {
      fmt.Println("Error reading file:", err)
      os.Exit(1)
   }
   _, err = io.Copy(cw, file)
   if err != nil {
      fmt.Println("Error copying:", err)
      os.Exit(1)
   }
   for c := range cw.Channel {
      fmt.printf("%c\n",c)
   }
}
}
