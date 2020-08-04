package process

import (
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

// NtSetInformationProcess is Unexported
var (
	ntdll                   = syscall.MustLoadDLL("ntdll.dll")
	ntSetInformationProcess = ntdll.MustFindProc("NtSetInformationProcess")
)

// Protect gets the current running process and initiates protection.
func Protect() {

	process, err := syscall.GetCurrentProcess()

	if err != nil {
		return
	}

	ntSetInformationProcess.Call(uintptr(process), 29, 1, 4)

	//fmt.Printf("testing")

}
