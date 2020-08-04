package Process

import "fmt"

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

	fmt.Printf("testing")

}
