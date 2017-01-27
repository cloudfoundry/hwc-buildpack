## hwc-buildpack

A Cloud Foundry [buildpack](http://docs.cloudfoundry.org/buildpacks/) for Windows applications.

Additional information can be found at [CloudFoundry.org](http://docs.cloudfoundry.org/buildpacks/).

## Dependencies
- [Golang Windows](https://golang.org/dl/)
- [Ginkgo](https://onsi.github.io/ginkgo/)
- Hostable Web Core
  - Install in Powershell by running `Install-WindowsFeature Web-WHC`

### Building

```
./scripts/build.sh

```

### Test

Unit Tests:

```
ginkgo -r -race
```


### Use in Cloud Foundry

Upload the buildpack to your Cloud Foundry and optionally specify it by name.

```
cf create-buildpack hwc_buildpack hwc_buildpack.zip 10
cf push my_app -b hwc_buildpack -s windows2012R2
```

### Help and Support

Join the #greenhouse channel in our [Slack community](http://slack.cloudfoundry.org/) if you need any further assistance.

### Active Development

The project backlog is on [Pivotal Tracker](https://www.pivotaltracker.com/n/projects/1156164).
