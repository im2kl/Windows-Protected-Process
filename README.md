## Windows Protected Process

Using the undocumented NtSetInformationProcess to gain process protection against kill.

http://undocumented.ntinternals.net/index.html?page=UserMode%2FUndocumented%20Functions%2FNT%20Objects%2FProcess%2FNtSetInformationProcess.html

After further testing the following will only protect against termination by non-admin tools, unelevated tools. tools with elevated permissions are able to terminate the process.