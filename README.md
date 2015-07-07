pogo
====
PowerShell as a Service


###Request Contexts
* __/command/__ - Executes named PowerShell command.
* __/script/__ - Executes named script.
* __/exit/__ - Terminates the service instance.

###Running Commands
Commands are handled by the /command/ context. Most common commands are readily supported and will return any structured data in JSON format. Parameters are passed via url query parameters. Named values will be broken out into key/value pairs and added to the commandstring.

Run __[Get-Date](https://technet.microsoft.com/en-us/library/hh849887.aspx)__ command.
    http://127.0.0.1:8080/command/Get-Date
Returns PowerShell Data in JSON  
```json
{
    "value":  "\/Date(1436237104231)\/",
    "DisplayHint":  2,
    "DateTime":  "Monday, July 6, 2015 9:45:04 PM"
}
```

Run __[Write-Host](https://technet.microsoft.com/en-us/library/ee177031.aspx)__ command.
    http://127.0.0.1:8080/command/Write-Host?"Hello, World!"
```json
Hello, World!
```

Run __[Get-Item](https://technet.microsoft.com/en-us/library/hh849788.aspx)__ command.
    http://127.0.0.1:8080/command/Get-Item?-Path="pogo.exe"
```json
{
    "Name":  "pogo.exe",
    "Length":  9076224,
    "DirectoryName":  "C:\\Users\\onyxhat\\Documents\\GitHub\\pogo",
    "Directory":  {
                      "Name":  "pogo",
                      "Parent":  {
                                     "Name":  "GitHub",
                                     "Parent":  "Documents",
                                     "Exists":  true,
                                     "Root":  "C:\\",
                                     "FullName":  "C:\\Users\\onyxhat\\Documents\\GitHub",
                                     "Extension":  "",
                                     "CreationTime":  "\/Date(1436236701491)\/",
                                     "CreationTimeUtc":  "\/Date(1436236701491)\/",
                                     "LastAccessTime":  "\/Date(1436236829685)\/",
                                     "LastAccessTimeUtc":  "\/Date(1436236829685)\/",
                                     "LastWriteTime":  "\/Date(1436236829685)\/",
                                     "LastWriteTimeUtc":  "\/Date(1436236829685)\/",
                                     "Attributes":  16
                                 },
                      "Exists":  true,
                      "Root":  {
                                   "Name":  "C:\\",
                                   "Parent":  null,
                                   "Exists":  true,
                                   "Root":  "C:\\",
                                   "FullName":  "C:\\",
                                   "Extension":  "",
                                   "CreationTime":  "\/Date(1426318734664)\/",
                                   "CreationTimeUtc":  "\/Date(1426318734664)\/",
                                   "LastAccessTime":  "\/Date(1436225518107)\/",
                                   "LastAccessTimeUtc":  "\/Date(1436225518107)\/",
                                   "LastWriteTime":  "\/Date(1436225518107)\/",
                                   "LastWriteTimeUtc":  "\/Date(1436225518107)\/",
                                   "Attributes":  22
                               },
                      "FullName":  "C:\\Users\\onyxhat\\Documents\\GitHub\\pogo",
                      "Extension":  "",
                      "CreationTime":  "\/Date(1436236829685)\/",
                      "CreationTimeUtc":  "\/Date(1436236829685)\/",
                      "LastAccessTime":  "\/Date(1436237018266)\/",
                      "LastAccessTimeUtc":  "\/Date(1436237018266)\/",
                      "LastWriteTime":  "\/Date(1436237018266)\/",
                      "LastWriteTimeUtc":  "\/Date(1436237018266)\/",
                      "Attributes":  16
                  },
    "IsReadOnly":  false,
    "Exists":  true,
    "FullName":  "C:\\Users\\onyxhat\\Documents\\GitHub\\pogo\\pogo.exe",
    "Extension":  ".exe",
    "CreationTime":  "\/Date(1436237018266)\/",
    "CreationTimeUtc":  "\/Date(1436237018266)\/",
    "LastAccessTime":  "\/Date(1436237018266)\/",
    "LastAccessTimeUtc":  "\/Date(1436237018266)\/",
    "LastWriteTime":  "\/Date(1436237020044)\/",
    "LastWriteTimeUtc":  "\/Date(1436237020044)\/",
    "Attributes":  32,
    "PSPath":  "Microsoft.PowerShell.Core\\FileSystem::C:\\Users\\onyxhat\\Documents\\GitHub\\pogo\\pogo.exe",
    "PSParentPath":  "Microsoft.PowerShell.Core\\FileSystem::C:\\Users\\onyxhat\\Documents\\GitHub\\pogo",
    "PSChildName":  "pogo.exe",
    "PSDrive":  {
                    "CurrentLocation":  "Users\\onyxhat\\Documents\\GitHub\\pogo",
                    "Name":  "C",
                    "Provider":  {
                                     "ImplementingType":  "Microsoft.PowerShell.Commands.FileSystemProvider",
                                     "HelpFile":  "System.Management.Automation.dll-Help.xml",
                                     "Name":  "FileSystem",
                                     "PSSnapIn":  "Microsoft.PowerShell.Core",
                                     "ModuleName":  "Microsoft.PowerShell.Core",
                                     "Module":  null,
                                     "Description":  "",
                                     "Capabilities":  52,
                                     "Home":  "C:\\Users\\onyxhat",
                                     "Drives":  "C M"
                                 },
                    "Root":  "C:\\",
                    "Description":  "",
                    "Credential":  {
                                       "UserName":  null,
                                       "Password":  null
                                   },
                    "DisplayRoot":  null
                },
    "PSProvider":  {
                       "ImplementingType":  {
                                                "Module":  "System.Management.Automation.dll",
                                                "Assembly":  "System.Management.Automation, Version=3.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35",
                                                "TypeHandle":  "System.RuntimeTypeHandle",
                                                "DeclaringMethod":  null,
                                                "BaseType":  "System.Management.Automation.Provider.NavigationCmdletProvider",
                                                "UnderlyingSystemType":  "Microsoft.PowerShell.Commands.FileSystemProvider",
                                                "FullName":  "Microsoft.PowerShell.Commands.FileSystemProvider",
                                                "AssemblyQualifiedName":  "Microsoft.PowerShell.Commands.FileSystemProvider, System.Management.Automation, Version=3.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35",
                                                "Namespace":  "Microsoft.PowerShell.Commands",
                                                "GUID":  "b4755d19-b6a7-38dc-ae06-4167f801062f",
                                                "IsEnum":  false,
                                                "GenericParameterAttributes":  null,
                                                "IsSecurityCritical":  true,
                                                "IsSecuritySafeCritical":  false,
                                                "IsSecurityTransparent":  false,
                                                "IsGenericTypeDefinition":  false,
                                                "IsGenericParameter":  false,
                                                "GenericParameterPosition":  null,
                                                "IsGenericType":  false,
                                                "IsConstructedGenericType":  false,
                                                "ContainsGenericParameters":  false,
                                                "StructLayoutAttribute":  "System.Runtime.InteropServices.StructLayoutAttribute",
                                                "Name":  "FileSystemProvider",
                                                "MemberType":  32,
                                                "DeclaringType":  null,
                                                "ReflectedType":  null,
                                                "MetadataToken":  33556309,
                                                "GenericTypeParameters":  "",
                                                "DeclaredConstructors":  "Void .ctor() Void .cctor()",
                                                "DeclaredEvents":  "",
                                                "DeclaredFields":  "System.Management.Automation.PSTraceSource tracer System.String ProviderName",
                                                "DeclaredMembers":  "System.String NormalizePath(System.String) System.IO.FileSystemInfo GetFileSystemInfo(System.String, Boolean ByRef) Boolean IsFilterSet() System.Object GetChildNamesDynamicParameters(System.String) System.Object GetChildItemsDynamicParameters(System.String, Boolean) System.String GetHelpMaml(System.String, System.String) System.Management.Automation.ProviderInfo Start(System.Management.Automation.ProviderInfo) System.Management.Automation.PSDriveInfo NewDrive(System.Management.Automation.PSDriveInfo) Void MapNetworkDrive(System.Management.Automation.PSDriveInfo) System.Management.Automation.PSDriveInfo RemoveDrive(System.Management.Automation.PSDriveInfo) Boolean IsSupportedDriveForPersistence(System.Management.Automation.PSDriveInfo) System.String GetUNCForNetworkDrive(System.String) System.Collections.ObjectModel.Collection`1[System.Management.Automation.PSDriveInfo] InitializeDefaultDrives() System.Object GetItemDynamicParameters(System.String) Boolean IsValidPath(System.String) Void GetItem(System.String) System.IO.FileSystemInfo GetFileSystemItem(System.String, Boolean ByRef, Boolean) Void InvokeDefaultAction(System.String) Void GetChildNames(System.String, System.Management.Automation.ReturnContainers) Boolean ConvertPath(System.String, System.String, System.String ByRef, System.String ByRef) System.Management.Automation.FlagsExpression`1[System.IO.FileAttributes] FormatAttributeSwitchParamters() System.String Mode(System.Management.Automation.PSObject) Void RenameItem(System.String, System.String) Void NewItem(System.String, System.String, System.Object) ItemType GetItemType(System.String) Void CreateDirectory(System.String, Boolean) Boolean CreateIntermediateDirectories(System.String) Void RemoveItem(System.String, Boolean) System.Object RemoveItemDynamicParameters(System.String, Boolean) Void RemoveDirectoryInfoItem(System.IO.DirectoryInfo, Boolean, Boolean, Boolean) Void RemoveFileInfoItem(System.IO.FileInfo, Boolean) Void RemoveFileSystemItem(System.IO.FileSystemInfo, Boolean) Boolean ItemExists(System.String) Boolean ItemExists(System.String, System.Management.Automation.ErrorRecord ByRef) System.Object ItemExistsDynamicParameters(System.String) Boolean HasChildItems(System.String) Boolean DirectoryInfoHasChildItems(System.IO.DirectoryInfo) Void CopyItem(System.String, System.String, Boolean) Void CopyDirectoryInfoItem(System.IO.DirectoryInfo, System.String, Boolean, Boolean) Void CopyFileInfoItem(System.IO.FileInfo, System.String, Boolean) System.String GetParentPath(System.String, System.String) Boolean IsAbsolutePath(System.String) Boolean IsUNCPath(System.String) Boolean IsUNCRoot(System.String) Boolean IsPathRoot(System.String) System.String NormalizeRelativePath(System.String, System.String) System.String NormalizeRelativePathHelper(System.String, System.String) System.String RemoveRelativeTokens(System.String) System.String GetCommonBase(System.String, System.String) System.Collections.Generic.Stack`1[System.String] TokenizePathToStack(System.String, System.String) System.Collections.Generic.Stack`1[System.String] NormalizeThePath(System.String, System.Collections.Generic.Stack`1[System.String]) System.String CreateNormalizedRelativePathFromStack(System.Collections.Generic.Stack`1[System.String]) System.String GetChildName(System.String) System.String EnsureDriveIsRooted(System.String) Boolean IsItemContainer(System.String) Void MoveItem(System.String, System.String) Void MoveFileInfoItem(System.IO.FileInfo, System.String, Boolean, Boolean) Void MoveDirectoryInfoItem(System.IO.DirectoryInfo, System.String, Boolean) Void CopyAndDelete(System.IO.DirectoryInfo, System.String, Boolean) Boolean IsSameVolume(System.String, System.String) Void GetProperty(System.String, System.Collections.ObjectModel.Collection`1[System.String]) System.Object GetPropertyDynamicParameters(System.String, System.Collections.ObjectModel.Collection`1[System.String]) Void SetProperty(System.String, System.Management.Automation.PSObject) System.Object SetPropertyDynamicParameters(System.String, System.Management.Automation.PSObject) Void ClearProperty(System.String, System.Collections.ObjectModel.Collection`1[System.String]) System.Object ClearPropertyDynamicParameters(System.String, System.Collections.ObjectModel.Collection`1[System.String]) System.Management.Automation.Provider.IContentReader GetContentReader(System.String) System.Object GetContentReaderDynamicParameters(System.String) System.Management.Automation.Provider.IContentWriter GetContentWriter(System.String) System.Object GetContentWriterDynamicParameters(System.String) Void ClearContent(System.String) System.Object ClearContentDynamicParameters(System.String) Int32 SafeGetFileAttributes(System.String) Void ValidateParameters(Boolean) Void GetSecurityDescriptor(System.String, System.Security.AccessControl.AccessControlSections) Void SetSecurityDescriptor(System.String, System.Security.AccessControl.ObjectSecurity) Void SetSecurityDescriptor(System.String, System.Security.AccessControl.ObjectSecurity, System.Security.AccessControl.AccessControlSections) System.Security.AccessControl.ObjectSecurity NewSecurityDescriptorFromPath(System.String, System.Security.AccessControl.AccessControlSections) System.Security.AccessControl.ObjectSecurity NewSecurityDescriptorOfType(System.String, System.Security.AccessControl.AccessControlSections) System.Security.AccessControl.ObjectSecurity NewSecurityDescriptor(ItemType) System.Management.Automation.ErrorRecord CreateErrorRecord(System.String, System.String) System.Object CopyItemDynamicParameters(System.String, System.String, Boolean) Boolean IsNetworkMappedDrive(System.Management.Automation.PSDriveInfo) System.String GetSubstitutedPathForNetworkDosDevice(System.String) System.String GetRootPathForNetworkDriveOrDosDevice(System.IO.DriveInfo) Void GetChildItems(System.String, Boolean, UInt32) Void GetPathItems(System.String, Boolean, UInt32, Boolean, System.Management.Automation.ReturnContainers) Void Dir(System.IO.DirectoryInfo, Boolean, UInt32, Boolean, System.Management.Automation.ReturnContainers) Boolean CheckItemExists(System.String, Boolean ByRef) Void CopyItemToRemoteSession(System.String, System.String, Boolean, Boolean, System.Management.Automation.Runspaces.PSSession) Void CopyItemFromRemoteSession(System.String, System.String, Boolean, Boolean, System.Management.Automation.Runspaces.PSSession) Void CopyItemLocal(System.String, System.String, Boolean, Boolean) Boolean CopyDirectoryFromRemoteSession(System.IO.DirectoryInfo, System.IO.DirectoryInfo, Boolean, Boolean, System.Management.Automation.PowerShell) Boolean RemoveFunctionPSCopyFileFromRemoteSession(System.Management.Automation.PowerShell) Boolean InitilizeFunctionPSCopyFileFromRemoteSession(System.Management.Automation.PowerShell) Boolean CopyFileFromRemoteSession(System.IO.FileInfo, System.IO.FileInfo, System.String, Boolean, System.Management.Automation.PowerShell, Int64) Boolean PerformCopyFileToRemoteSession(System.IO.FileInfo, System.String, System.Management.Automation.PowerShell) Boolean CopyFileToRemoteSession(System.IO.FileInfo, System.String, Boolean, System.Management.Automation.PowerShell) Boolean CopyDirectoryToRemoteSession(System.IO.DirectoryInfo, System.String, Boolean, Boolean, System.Management.Automation.PowerShell) Boolean RemoteDestinationPathIsFile(System.String, System.Management.Automation.PowerShell) System.String CreateDirectoryOnRemoteSession(System.IO.DirectoryInfo, System.String, Boolean, System.Management.Automation.PowerShell) Void .ctor() Void .cctor() System.Management.Automation.PSTraceSource tracer System.String ProviderName Microsoft.PowerShell.Commands.FileSystemProvider+ItemType Microsoft.PowerShell.Commands.FileSystemProvider+NativeMethods Microsoft.PowerShell.Commands.FileSystemProvider+NetResource",
                                                "DeclaredMethods":  "System.String Mode(System.Management.Automation.PSObject) System.String NormalizePath(System.String) System.IO.FileSystemInfo GetFileSystemInfo(System.String, Boolean ByRef) Boolean IsFilterSet() System.Object GetChildNamesDynamicParameters(System.String) System.Object GetChildItemsDynamicParameters(System.String, Boolean) System.String GetHelpMaml(System.String, System.String) System.Management.Automation.ProviderInfo Start(System.Management.Automation.ProviderInfo) System.Management.Automation.PSDriveInfo NewDrive(System.Management.Automation.PSDriveInfo) Void MapNetworkDrive(System.Management.Automation.PSDriveInfo) System.Management.Automation.PSDriveInfo RemoveDrive(System.Management.Automation.PSDriveInfo) Boolean IsSupportedDriveForPersistence(System.Management.Automation.PSDriveInfo) System.String GetUNCForNetworkDrive(System.String) System.Collections.ObjectModel.Collection`1[System.Management.Automation.PSDriveInfo] InitializeDefaultDrives() System.Object GetItemDynamicParameters(System.String) Boolean IsValidPath(System.String) Void GetItem(System.String) System.IO.FileSystemInfo GetFileSystemItem(System.String, Boolean ByRef, Boolean) Void InvokeDefaultAction(System.String) Void GetChildNames(System.String, System.Management.Automation.ReturnContainers) Boolean ConvertPath(System.String, System.String, System.String ByRef, System.String ByRef) System.Management.Automation.FlagsExpression`1[System.IO.FileAttributes] FormatAttributeSwitchParamters() Void RenameItem(System.String, System.String) Void NewItem(System.String, System.String, System.Object) ItemType GetItemType(System.String) Void CreateDirectory(System.String, Boolean) Boolean CreateIntermediateDirectories(System.String) Void RemoveItem(System.String, Boolean) System.Object RemoveItemDynamicParameters(System.String, Boolean) Void RemoveDirectoryInfoItem(System.IO.DirectoryInfo, Boolean, Boolean, Boolean) Void RemoveFileInfoItem(System.IO.FileInfo, Boolean) Void RemoveFileSystemItem(System.IO.FileSystemInfo, Boolean) Boolean ItemExists(System.String) Boolean ItemExists(System.String, System.Management.Automation.ErrorRecord ByRef) System.Object ItemExistsDynamicParameters(System.String) Boolean HasChildItems(System.String) Boolean DirectoryInfoHasChildItems(System.IO.DirectoryInfo) Void CopyItem(System.String, System.String, Boolean) Void CopyDirectoryInfoItem(System.IO.DirectoryInfo, System.String, Boolean, Boolean) Void CopyFileInfoItem(System.IO.FileInfo, System.String, Boolean) System.String GetParentPath(System.String, System.String) Boolean IsAbsolutePath(System.String) Boolean IsUNCPath(System.String) Boolean IsUNCRoot(System.String) Boolean IsPathRoot(System.String) System.String NormalizeRelativePath(System.String, System.String) System.String NormalizeRelativePathHelper(System.String, System.String) System.String RemoveRelativeTokens(System.String) System.String GetCommonBase(System.String, System.String) System.Collections.Generic.Stack`1[System.String] TokenizePathToStack(System.String, System.String) System.Collections.Generic.Stack`1[System.String] NormalizeThePath(System.String, System.Collections.Generic.Stack`1[System.String]) System.String CreateNormalizedRelativePathFromStack(System.Collections.Generic.Stack`1[System.String]) System.String GetChildName(System.String) System.String EnsureDriveIsRooted(System.String) Boolean IsItemContainer(System.String) Void MoveItem(System.String, System.String) Void MoveFileInfoItem(System.IO.FileInfo, System.String, Boolean, Boolean) Void MoveDirectoryInfoItem(System.IO.DirectoryInfo, System.String, Boolean) Void CopyAndDelete(System.IO.DirectoryInfo, System.String, Boolean) Boolean IsSameVolume(System.String, System.String) Void GetProperty(System.String, System.Collections.ObjectModel.Collection`1[System.String]) System.Object GetPropertyDynamicParameters(System.String, System.Collections.ObjectModel.Collection`1[System.String]) Void SetProperty(System.String, System.Management.Automation.PSObject) System.Object SetPropertyDynamicParameters(System.String, System.Management.Automation.PSObject) Void ClearProperty(System.String, System.Collections.ObjectModel.Collection`1[System.String]) System.Object ClearPropertyDynamicParameters(System.String, System.Collections.ObjectModel.Collection`1[System.String]) System.Management.Automation.Provider.IContentReader GetContentReader(System.String) System.Object GetContentReaderDynamicParameters(System.String) System.Management.Automation.Provider.IContentWriter GetContentWriter(System.String) System.Object GetContentWriterDynamicParameters(System.String) Void ClearContent(System.String) System.Object ClearContentDynamicParameters(System.String) Int32 SafeGetFileAttributes(System.String) Void ValidateParameters(Boolean) Void GetSecurityDescriptor(System.String, System.Security.AccessControl.AccessControlSections) Void SetSecurityDescriptor(System.String, System.Security.AccessControl.ObjectSecurity) Void SetSecurityDescriptor(System.String, System.Security.AccessControl.ObjectSecurity, System.Security.AccessControl.AccessControlSections) System.Security.AccessControl.ObjectSecurity NewSecurityDescriptorFromPath(System.String, System.Security.AccessControl.AccessControlSections) System.Security.AccessControl.ObjectSecurity NewSecurityDescriptorOfType(System.String, System.Security.AccessControl.AccessControlSections) System.Security.AccessControl.ObjectSecurity NewSecurityDescriptor(ItemType) System.Management.Automation.ErrorRecord CreateErrorRecord(System.String, System.String) System.Object CopyItemDynamicParameters(System.String, System.String, Boolean) Boolean IsNetworkMappedDrive(System.Management.Automation.PSDriveInfo) System.String GetSubstitutedPathForNetworkDosDevice(System.String) System.String GetRootPathForNetworkDriveOrDosDevice(System.IO.DriveInfo) Void GetChildItems(System.String, Boolean, UInt32) Void GetPathItems(System.String, Boolean, UInt32, Boolean, System.Management.Automation.ReturnContainers) Void Dir(System.IO.DirectoryInfo, Boolean, UInt32, Boolean, System.Management.Automation.ReturnContainers) Boolean CheckItemExists(System.String, Boolean ByRef) Void CopyItemToRemoteSession(System.String, System.String, Boolean, Boolean, System.Management.Automation.Runspaces.PSSession) Void CopyItemFromRemoteSession(System.String, System.String, Boolean, Boolean, System.Management.Automation.Runspaces.PSSession) Void CopyItemLocal(System.String, System.String, Boolean, Boolean) Boolean CopyDirectoryFromRemoteSession(System.IO.DirectoryInfo, System.IO.DirectoryInfo, Boolean, Boolean, System.Management.Automation.PowerShell) Boolean RemoveFunctionPSCopyFileFromRemoteSession(System.Management.Automation.PowerShell) Boolean InitilizeFunctionPSCopyFileFromRemoteSession(System.Management.Automation.PowerShell) Boolean CopyFileFromRemoteSession(System.IO.FileInfo, System.IO.FileInfo, System.String, Boolean, System.Management.Automation.PowerShell, Int64) Boolean PerformCopyFileToRemoteSession(System.IO.FileInfo, System.String, System.Management.Automation.PowerShell) Boolean CopyFileToRemoteSession(System.IO.FileInfo, System.String, Boolean, System.Management.Automation.PowerShell) Boolean CopyDirectoryToRemoteSession(System.IO.DirectoryInfo, System.String, Boolean, Boolean, System.Management.Automation.PowerShell) Boolean RemoteDestinationPathIsFile(System.String, System.Management.Automation.PowerShell) System.String CreateDirectoryOnRemoteSession(System.IO.DirectoryInfo, System.String, Boolean, System.Management.Automation.PowerShell)",
                                                "DeclaredNestedTypes":  "Microsoft.PowerShell.Commands.FileSystemProvider+ItemType Microsoft.PowerShell.Commands.FileSystemProvider+NativeMethods Microsoft.PowerShell.Commands.FileSystemProvider+NetResource",
                                                "DeclaredProperties":  "",
                                                "ImplementedInterfaces":  "System.Management.Automation.IResourceSupplier System.Management.Automation.Provider.IContentCmdletProvider System.Management.Automation.Provider.IPropertyCmdletProvider System.Management.Automation.Provider.ISecurityDescriptorCmdletProvider System.Management.Automation.Provider.ICmdletProviderSupportsHelp",
                                                "TypeInitializer":  "Void .cctor()",
                                                "IsNested":  false,
                                                "Attributes":  1048833,
                                                "IsVisible":  true,
                                                "IsNotPublic":  false,
                                                "IsPublic":  true,
                                                "IsNestedPublic":  false,
                                                "IsNestedPrivate":  false,
                                                "IsNestedFamily":  false,
                                                "IsNestedAssembly":  false,
                                                "IsNestedFamANDAssem":  false,
                                                "IsNestedFamORAssem":  false,
                                                "IsAutoLayout":  true,
                                                "IsLayoutSequential":  false,
                                                "IsExplicitLayout":  false,
                                                "IsClass":  true,
                                                "IsInterface":  false,
                                                "IsValueType":  false,
                                                "IsAbstract":  false,
                                                "IsSealed":  true,
                                                "IsSpecialName":  false,
                                                "IsImport":  false,
                                                "IsSerializable":  false,
                                                "IsAnsiClass":  true,
                                                "IsUnicodeClass":  false,
                                                "IsAutoClass":  false,
                                                "IsArray":  false,
                                                "IsByRef":  false,
                                                "IsPointer":  false,
                                                "IsPrimitive":  false,
                                                "IsCOMObject":  false,
                                                "HasElementType":  false,
                                                "IsContextful":  false,
                                                "IsMarshalByRef":  false,
                                                "GenericTypeArguments":  "",
                                                "CustomAttributes":  "[System.Management.Automation.OutputTypeAttribute(new Type[2] { typeof(System.String), typeof(System.IO.FileInfo) }, ProviderCmdlet = \"New-Item\")] [System.Management.Automation.OutputTypeAttribute(new Type[1] { typeof(System.IO.FileInfo) }, ProviderCmdlet = \"Get-Item\")] [System.Management.Automation.OutputTypeAttribute(new Type[1] { typeof(System.Management.Automation.PathInfo) }, ProviderCmdlet = \"Push-Location\")] [System.Management.Automation.OutputTypeAttribute(new Type[5] { typeof(System.Boolean), typeof(System.String), typeof(System.DateTime), typeof(System.IO.FileInfo), typeof(System.IO.DirectoryInfo) }, ProviderCmdlet = \"Get-ItemProperty\")] [System.Management.Automation.OutputTypeAttribute(new Type[1] { typeof(System.Security.AccessControl.FileSecurity) }, ProviderCmdlet = \"Set-Acl\")] [System.Management.Automation.OutputTypeAttribute(new Type[2] { typeof(System.String), typeof(System.Management.Automation.PathInfo) }, ProviderCmdlet = \"Resolve-Path\")] [System.Management.Automation.OutputTypeAttribute(new Type[2] { typeof(System.Byte), typeof(System.String) }, ProviderCmdlet = \"Get-Content\")] [System.Management.Automation.Provider.CmdletProviderAttribute(\"FileSystem\", (System.Management.Automation.Provider.ProviderCapabilities)52)] [System.Management.Automation.OutputTypeAttribute(new Type[2] { typeof(System.IO.FileInfo), typeof(System.IO.DirectoryInfo) }, ProviderCmdlet = \"Get-ChildItem\")] [System.Management.Automation.OutputTypeAttribute(new Type[2] { typeof(System.Security.AccessControl.FileSecurity), typeof(System.Security.AccessControl.DirectorySecurity) }, ProviderCmdlet = \"Get-Acl\")] [System.Management.Automation.OutputTypeAttribute(new Type[4] { typeof(System.Boolean), typeof(System.String), typeof(System.IO.FileInfo), typeof(System.IO.DirectoryInfo) }, ProviderCmdlet = \"Get-Item\")]"
                                            },
                       "HelpFile":  "System.Management.Automation.dll-Help.xml",
                       "Name":  "FileSystem",
                       "PSSnapIn":  {
                                        "Name":  "Microsoft.PowerShell.Core",
                                        "IsDefault":  true,
                                        "ApplicationBase":  "C:\\Windows\\System32\\WindowsPowerShell\\v1.0",
                                        "AssemblyName":  "System.Management.Automation, Version=3.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35, ProcessorArchitecture=MSIL",
                                        "ModuleName":  "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\System.Management.Automation.dll",
                                        "PSVersion":  "5.0.10162.0",
                                        "Version":  "3.0.0.0",
                                        "Types":  "types.ps1xml typesv3.ps1xml",
                                        "Formats":  "Certificate.format.ps1xml DotNetTypes.format.ps1xml FileSystem.format.ps1xml Help.format.ps1xml HelpV3.format.ps1xml PowerShellCore.format.ps1xml PowerShellTrace.format.ps1xml Registry.format.ps1xml",
                                        "Description":  "This Windows PowerShell snap-in contains cmdlets used to manage components of Windows PowerShell.",
                                        "Vendor":  "Microsoft Corporation",
                                        "LogPipelineExecutionDetails":  false
                                    },
                       "ModuleName":  "Microsoft.PowerShell.Core",
                       "Module":  null,
                       "Description":  "",
                       "Capabilities":  52,
                       "Home":  "C:\\Users\\onyxhat",
                       "Drives":  [
                                      "C",
                                      "M"
                                  ]
                   },
    "PSIsContainer":  false,
    "VersionInfo":  {
                        "Comments":  null,
                        "CompanyName":  null,
                        "FileBuildPart":  0,
                        "FileDescription":  null,
                        "FileMajorPart":  0,
                        "FileMinorPart":  0,
                        "FileName":  "C:\\Users\\onyxhat\\Documents\\GitHub\\pogo\\pogo.exe",
                        "FilePrivatePart":  0,
                        "FileVersion":  null,
                        "InternalName":  null,
                        "IsDebug":  false,
                        "IsPatched":  false,
                        "IsPrivateBuild":  false,
                        "IsPreRelease":  false,
                        "IsSpecialBuild":  false,
                        "Language":  null,
                        "LegalCopyright":  null,
                        "LegalTrademarks":  null,
                        "OriginalFilename":  null,
                        "PrivateBuild":  null,
                        "ProductBuildPart":  0,
                        "ProductMajorPart":  0,
                        "ProductMinorPart":  0,
                        "ProductName":  null,
                        "ProductPrivatePart":  0,
                        "ProductVersion":  null,
                        "SpecialBuild":  null,
                        "FileVersionRaw":  {
                                               "Major":  0,
                                               "Minor":  0,
                                               "Build":  0,
                                               "Revision":  0,
                                               "MajorRevision":  0,
                                               "MinorRevision":  0
                                           },
                        "ProductVersionRaw":  {
                                                  "Major":  0,
                                                  "Minor":  0,
                                                  "Build":  0,
                                                  "Revision":  0,
                                                  "MajorRevision":  0,
                                                  "MinorRevision":  0
                                              }
                    },
    "BaseName":  "pogo",
    "Target":  {

               },
    "LinkType":  null,
    "Mode":  "-a----"
}
```

###TODO
* Add ___Authentication___
* Implement POST method of form values as parameters (takes precedence over query params)
* Command restrictions
* Additional configuration values
* Remote configuration
