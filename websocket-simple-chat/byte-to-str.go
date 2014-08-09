package main

import(
    "bytes"
)

func ByteToStr(b []byte) string {

    buffer := bytes.NewBuffer(b)

    return buffer.String()
}
