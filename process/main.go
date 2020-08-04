package process

import (
	"fmt"
	"syscall"
)

/*
NTSYSAPI
NTSTATUS
NTAPI

NtSetInformationProcess(
  IN HANDLE               ProcessHandle,
  IN PROCESS_INFORMATION_CLASS ProcessInformationClass,
  IN PVOID                ProcessInformation,
  IN ULONG                ProcessInformationLength );
);

*/

func Protect() {

	_, err := syscall.GetCurrentProcess()

	if err != nil {
		return
	}
	fmt.Printf("testing")

}
