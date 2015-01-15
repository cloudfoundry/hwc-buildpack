Configuring HTTP.SYS
------------------------------------------------------
http://msdn.microsoft.com/en-us/library/ms733768.aspx
netsh http add urlacl url=http://foobar:port/ user=containerUser
netsh http delete urlacl url=http://foobar:port/ user=containerUser

Hostable Web Core
------------------------------------------------------
API Ref: http://msdn.microsoft.com/en-us/library/ms693832(v=vs.90).aspx
Walkthrough: http://msdn.microsoft.com/en-us/library/ms689327(v=vs.90).aspx
Runtime errors show up in windows application event log, source = HostableWebCore

Note that WebcoreActivate and WebcoreShutdown can be called only once in process lifetime.
Also HWC always uses .Net 2.0 schema and there is no way to make it pick .Net 4.0 schema.
So root web.config passed to WebcoreActivate cannot have properties which are only supported
 in .Net 4.0. This problem will be fixed in Win8.

 Notes:
 -----------------------------------------------------
 - On Windows 7, you can only use .net 2.0 runtime version
(read above for explanation, this is hostable web core limitation).
 
 given the following args:
  -p 8869 -r d:\container -v 2.0 -b true

  It expects your web app to live in: d:\container\www
  It will create log in d:\container\log
  It will create config in d:\container\config
  It will create temp in d:\container\tmp