hwc-buildpack
=============

A Cloud Foundry [buildpack](http://docs.cloudfoundry.org/buildpacks/) for
Windows applications.

Additional information can be found at
[CloudFoundry.org](http://docs.cloudfoundry.org/buildpacks/).

## Dependencies
- 64 bit version of Windows (tested with Windows Server 2012 R2 Standard)
- msbuild in PATH
- Administrator access

Building in Visual Studio
========================

1. Install https://visualstudiogallery.msdn.microsoft.com/7a52473f-9e1a-40f3-8bd8-6c00ab163609 (nspec test runner)

1. Open Visual Studio as Administrator.

1. Use in Cloud Foundry

Upload the buildpack to your Cloud Foundry and optionally specify it by name.

```
cf create-buildpack hwc-buildpack hwc-buildpack.zip 10
cf push my_app -b hwc-buildpack -s windows2012R2
```

Help and Support
================

Join the #greenhouse channel in our [Slack
community](http://slack.cloudfoundry.org/) if you need any further assistance.

Active Development
=================

The project backlog is on [Pivotal
Tracker](https://www.pivotaltracker.com/n/projects/1156164).
