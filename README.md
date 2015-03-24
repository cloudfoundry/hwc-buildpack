windows app lifecycle 
=====================

The windows application lifecycle implements the traditional Cloud Foundry deployment
strategy.

The **Builder** downloads buildpacks and app bits, and produces a droplet.

The **Launcher** runs the start command using a standard rootfs and
environment.

The **Healthcheck** runs a tcp port check, defaulting to port 8080.

The **WebAppServer** runs a HostableWebCore server to host the user's app.

Read about the app lifecycle spec here: https://github.com/cloudfoundry-incubator/diego-design-notes#app-lifecycles

## Dependencies
- 64 bit version of Windows (tested with Windows Server 2012 R2 Standard)
- msbuild in PATH
- Administrator access

building on the command line
============================

1. Run make.bat in cmd.

building in Visual Studio
========================

1. Install https://visualstudiogallery.msdn.microsoft.com/7a52473f-9e1a-40f3-8bd8-6c00ab163609 (nspec test runner)

1. Open Visual Studio as Administrator.
![opening as admin](https://github.com/pivotal-cf-experimental/containerizer/blob/master/README_images/open_as_admin.png)
![visual studio running as admin](https://github.com/pivotal-cf-experimental/containerizer/blob/master/README_images/showing_vs_running_as_admin.png)
