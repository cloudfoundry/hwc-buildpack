## hwc-buildpack

A Cloud Foundry [buildpack](http://docs.cloudfoundry.org/buildpacks/) for Windows applications.

Additional information can be found at [CloudFoundry.org](http://docs.cloudfoundry.org/buildpacks/).

## Dependencies
- [Golang Windows](https://golang.org/dl/)
- [Ginkgo](https://onsi.github.io/ginkgo/)
- [Hostable Web Core](https://github.com/cloudfoundry-incubator/hwc)

### Building

```
./scripts/build.sh

```

### Test

Unit Tests:

```
cd src/compile
ginkgo -r -race
```

Integration Tests (must be run against a Cloud Foundry deployment with Windows cells):

```
BUNDLE_GEMFILE=cf.Gemfile bundle exec buildpack-build
```

### Use in Cloud Foundry

Upload the buildpack to your Cloud Foundry and optionally specify it by name.

```
BUNDLE_GEMFILE=cf.Gemfile bundle exec buildpack-packager --{cached | uncached}
cf create-buildpack hwc_buildpack hwc_buildpack-<cache/version info>.zip 10
cf push my_app -b hwc_buildpack -s windows2012R2
```

### Help and Support

Join the #greenhouse channel in our [Slack community](http://slack.cloudfoundry.org/) if you need any further assistance.

### Active Development

The project backlog is on [Pivotal Tracker](https://www.pivotaltracker.com/n/projects/1042066).
