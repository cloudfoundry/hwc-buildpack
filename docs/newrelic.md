# HWC Buildpack with New Relic Agent Support
This buildpack has support for New Relic .NET Framework agent.

## How the support for New Relic Agent works

The new relic logic is executed only if one of the following conditions is true:
* **"NEW_RELIC_LICENSE_KEY"** exists
* **"NEW_RELIC_DOWNLOAD_URL"** exists
* there is a user-provided-service with the word **"newrelic"** as part of the name
* there is a SERVICE in VCAP_SERVICES with the name **"newrelic"**
* for cached buildpack: in the buildpack's manifest.yml file there is a **dependency entity** for newrelic with its **file** property set to the name of the cached version of New Relic agent


You can bind your application to New Relic in the following ways:
* using environment variables
* using New Relic's Agent Tile
* using a user-provided-service

The buildpack contains a copy of **newrelic.config** file with minimum requirements to successfully bind to an application. If you need to customize New Relic agent's configuration, you need to provide your own **newrelic.config** file by copying it to your application folder before pushing the application. If a copy of **newrelic.config** exists in the application folder, the buildpack uses that.

New Relic agent requires several environment variables to be set in order for the agent to be loaded up by **"hwc.exe"**. For this reason, the buildpack contains a **"Procfile"** which invokes **"run.cmd"** command file. The **"run.cmd"** command file is built by the buildpack inside of the application folder, and contains the following environment variables:

    set COR_ENABLE_PROFILING=1
    set COR_PROFILER={71DA0A04-7777-4EC6-9643-7D28B46A8A41}
    set COR_PROFILER_PATH=%~dp0newrelic\NewRelic.Profiler.dll
    set NEWRELIC_HOME=%~dp0newrelic
    set NEWRELIC_INSTALL_PATH=%~dp0newrelic

plus, at the end it invokes **".cloudfoundry\hwc.exe"**.

**Note:** If the buildpack finds a **Procfile** in the application folder, it does not use its own copy of **Procfile**. So in order for New Relic agent to work properly, you need to make sure you set the above environment variables, to make profiler path and other values available to New Relic agent.


### Using New Relic Agent

You need to specify a New Relic account **license key** in one of the following ways in order to bind your application to New Relic service:

* Bind your application to New Relic using the Agent Tile in the Marketplace
	- click on New Relic tile in the Marketplace
	- create an instance of the service for the plan that is associated with your target New Relic account
	- use "services" section in your manifest.yml file, and specify the name of the service you just created in the marketplace
	- restage the application using **"cf restage YOUR_APPNAME"**

* Bind your application to New Relic using AppMgr
	- Push your application to PCF
	- In AppMgr click on your application
	- goto "Service" tab of your application
	- If you have already created a service instance from the tile, select "Bind Service". If you have not created any service instances, and this is the first time, select "New Service"
	- Follow the instructions to create a new service or bind to an existing service
	- "restage" the application using **"cf restage YOUR_APPNAME"**

* Bind your application to New Relic using User-Provided-Service
	- Create 1 user-provided-service with the word "newrelic" embedded as part of the service name 
	- add the following credentials to the user-rpovided-service:
		- "licenseKey" This is New Relic License Key - **REQUIRED**
		- "appName"    If you want to change the app name in New Relic use this property - **OPTIONAL**
	- push your application in one of the following ways:
		- by adding the user-provided-service to the application manifest.yml before pushing the app
		- by adding the user-provided-service in AppMgr and restaging it after you add the service


### Using Environment Variables or "newrelic.config" file

* You can alternatively use combination of "newrelic.config" file and/or environment variables to configure New Relic dotNet agent to report your application's health and performance to the designated New Relic account.

	- A copy of the 'newrelic.config' file is provided with the buildpack. If you need to add any agent features such as proxy settings, or change any other agent settings such as logging behavior, copy **newrelic.config** file from the agent folder into the application folder, and edit as required. The following are some examplles you can use:

		- add your New Relic license key:
			```
			  <service licenseKey="9999999999999999999999999999999999999999">
			```
		
			alternatively you can add the license key to application's 'manifest.yml' file as an environment variable "NEW_RELIC_LICENSE_KEY" in the "env" section


		- add the New Relic application name as you'd like it to appear in New Relic
			```
			  <application>
			    <name>My Application</name>
			  </application>
			```

			alternatively you can add the New Relic app name to application's 'manifest.yml' file as an environment variable "NEW_RELIC_APP_NAME" in the "env" section


	    - add proxy settings to the "service" element as a sub-element. example:
			```
	    	  <service licenseKey="9999999999999999999999999999999999999999">
	    	    <proxy host="my_proxy_server.com" port="9090" />
	    	  </service>
			```

	    - change agent logging level and destination
			```
	    	  <log level="info" console="true" />
			```

	    - as 'hwc.exe' is the executable running your application, make sure 'newrelic.config' contains the following tag:
			```
			  <instrumentation>
			    <applications>
			      <application name="hwc.exe" />
			    </applications>
			  </instrumentation>
			```

			this section is required. So if you plan to use your own copy of **newrelic.config** make sure that it contains this section

	    
	    Note:  Depending on your CI/CD pipeline, the Application directory may be created on-the-fly as part of the pipeline.  If that is the case and you are modifying this file, your pipeline will need to copy over the file to the Application directory before deploying/pushing the app to PCF.

	 
	- Push your application to PCF using this buildpack. To do that, edit your manifest.yml and add/update the following entry.

		buildpack: hwc_buildpack

		Then run "cf push".

		Note: If this is CI/CD (aka Bamboo), the "cf push" may not be required as your pipeline internally uses "cf push" to push the application to PCF.


* Check the logs. 

	Use 'cf logs <APP_NAME>' or 'cf logs <APP_NAME> --recent' to examine the logs. It should display New Relic agent installation progress.

### Cached Buildpack

The cached version of hwc buildpack contains New Relic agent. If you use the cached version by default the buildpack uses the embedded version of the agent. 

**Note:** In order to use a different version of New Relic agent in **disconnected environments**, you can download newrelic agent from http://download.newrelic.com/dot_net_agent/previous_releases/newrelic-agent-win-x64-8.2.216.0.zip (or the latest **"zipped"** version of the agent) and host it locally in your network. Then specify your internal location as an ENV variable - **"NEW_RELIC_DOWNLOAD_URL"** in your application **manifest.yml** file.




