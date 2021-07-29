// Package services provides all services that adapters can use to work.
package services

import (
	"github.com/wy7-source/iti-challenge/application/domain"
	"github.com/wy7-source/iti-challenge/application/services/validators"
)

// PasswordService is the principal service.
type PasswordService struct {
}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

// Validate is used to validate the password.
// Receive the Password and validate it with multithreading.
// It returns if the password is valid, or not.
func (p *PasswordService) Validate(password domain.PasswordInterface) bool {
	/*
		For performance purpouses only, we gonna work with channels;
		"Channels are a typed conduit through which you can send and receive values with the channel operator,<-." - https://tour.golang.org/concurrency/2

		With channels, we can separate each validation in one go routine, making a multithreading validation, and share or read some data between routines.

		To read all validation results, we need to use a loop with range. But while the channel are open, the loop wont finish, generating a infinite loop.

		To Prevent this, we use a Semaphore Pattern in Go.

		The semaphore is used for routines to notify when they complete.

		And with this, we can closes the Validation Channel  with value of validations, and auto closes the channel of semaphores too.
	*/

	// Channels
	validationChan := make(chan bool)
	finishedChan := make(chan bool)

	// Routines
	go func() {
		validators.DigitValidator(password.GetValue(), validationChan)
		finishedChan <- true
	}()
	go func() {
		validators.DuplicatesValidator(password.GetValue(), validationChan)
		finishedChan <- true
	}()
	go func() {
		validators.LowerValidator(password.GetValue(), validationChan)
		finishedChan <- true
	}()
	go func() {
		validators.SpecialCharsValidator(password.GetValue(), validationChan)
		finishedChan <- true
	}()
	go func() {
		validators.UpperValidator(password.GetValue(), validationChan)
		finishedChan <- true
	}()
	go func() {
		validators.WhiteSpacesValidator(password.GetValue(), validationChan)
		finishedChan <- true
	}()
	go func() {
		validators.MinLenghtValidator(password.GetValue(), validationChan)
		finishedChan <- true
	}()

	// Semaphore
	go func() {
		<-finishedChan
		<-finishedChan
		<-finishedChan
		<-finishedChan
		<-finishedChan
		<-finishedChan
		<-finishedChan
		close(validationChan)
		close(finishedChan)
	}()
	// If some routine return false, we immediatly return false.
	for isValid := range validationChan {
		if !isValid {
			return false
		}
	}
	// If all routines return true, the password is valid and we return true.
	return true
}
