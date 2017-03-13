package main

type AuthACK struct {
	Message

	MPMMajorVersion uint8
	MPMMinorVersion uint8
}
