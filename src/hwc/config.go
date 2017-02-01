package main

const AspnetConfig = `<?xml version="1.0" encoding="UTF-8" ?>
<configuration>
  <runtime>
    <legacyUnhandledExceptionPolicy enabled="false" />
    <legacyImpersonationPolicy enabled="true"/>
    <alwaysFlowImpersonationPolicy enabled="false"/>
    <SymbolReadingPolicy enabled="1" />
  </runtime>
  <startup useLegacyV2RuntimeActivationPolicy="true" />
</configuration>
`
const ApplicationHostConfig = `<?xml version="1.0" encoding="UTF-8"?>
<configuration>
  <configSections>
    <sectionGroup name="system.applicationHost">
      <section name="applicationPools" allowDefinition="AppHostOnly" overrideModeDefault="Deny" />
      <section name="configHistory" allowDefinition="AppHostOnly" overrideModeDefault="Deny" />
      <section name="customMetadata" allowDefinition="AppHostOnly" overrideModeDefault="Deny" />
      <section name="listenerAdapters" allowDefinition="AppHostOnly" overrideModeDefault="Deny" />
      <section name="log" allowDefinition="AppHostOnly" overrideModeDefault="Deny" />
      <section name="serviceAutoStartProviders" allowDefinition="AppHostOnly" overrideModeDefault="Deny" />
      <section name="sites" allowDefinition="AppHostOnly" overrideModeDefault="Deny" />
      <section name="webLimits" allowDefinition="AppHostOnly" overrideModeDefault="Deny" />
    </sectionGroup>
    <sectionGroup name="system.webServer">
      <section name="asp" overrideModeDefault="Deny" />
      <section name="caching" overrideModeDefault="Allow" />
      <section name="cgi" overrideModeDefault="Deny" />
      <section name="defaultDocument" overrideModeDefault="Allow" />
      <section name="directoryBrowse" overrideModeDefault="Allow" />
      <section name="fastCgi" allowDefinition="AppHostOnly" overrideModeDefault="Deny" />
      <section name="globalModules" allowDefinition="AppHostOnly" overrideModeDefault="Deny" />
      <section name="handlers" overrideModeDefault="Allow" />
      <section name="httpCompression" allowDefinition="AppHostOnly" overrideModeDefault="Deny" />
      <section name="httpErrors" overrideModeDefault="Allow" />
      <section name="httpLogging" overrideModeDefault="Deny" />
      <section name="httpProtocol" overrideModeDefault="Allow" />
      <section name="httpRedirect" overrideModeDefault="Allow" />
      <section name="httpTracing" overrideModeDefault="Allow" />
      <section name="isapiFilters" allowDefinition="MachineToApplication" overrideModeDefault="Deny" />
      <section name="modules" allowDefinition="MachineToApplication" overrideModeDefault="Allow" />
      <section name="odbcLogging" overrideModeDefault="Deny" />
      <sectionGroup name="security">
        <section name="access" overrideModeDefault="Deny" />
        <section name="applicationDependencies" overrideModeDefault="Deny" />
        <sectionGroup name="authentication">
          <section name="anonymousAuthentication" overrideModeDefault="Deny" />
          <section name="basicAuthentication" overrideModeDefault="Deny" />
          <section name="clientCertificateMappingAuthentication" overrideModeDefault="Deny" />
          <section name="digestAuthentication" overrideModeDefault="Deny" />
          <section name="iisClientCertificateMappingAuthentication" overrideModeDefault="Deny" />
          <section name="windowsAuthentication" overrideModeDefault="Allow" />
        </sectionGroup>
        <section name="authorization" overrideModeDefault="Allow" />
        <section name="ipSecurity" overrideModeDefault="Deny" />
        <section name="isapiCgiRestriction" allowDefinition="AppHostOnly" overrideModeDefault="Deny" />
        <section name="requestFiltering" overrideModeDefault="Allow" />
      </sectionGroup>
      <section name="serverRuntime" overrideModeDefault="Deny" />
      <section name="serverSideInclude" overrideModeDefault="Deny" />
      <section name="staticContent" overrideModeDefault="Allow" />
      <sectionGroup name="tracing">
        <section name="traceFailedRequests" overrideModeDefault="Allow" />
        <section name="traceProviderDefinitions" overrideModeDefault="Allow" />
      </sectionGroup>
      <section name="urlCompression" overrideModeDefault="Allow" />
      <section name="validation" overrideModeDefault="Allow" />
      <sectionGroup name="webdav">
        <section name="globalSettings" overrideModeDefault="Deny" />
        <section name="authoring" overrideModeDefault="Deny" />
        <section name="authoringRules" overrideModeDefault="Deny" />
      </sectionGroup>
      <sectionGroup name="wdeploy">
        <section name="backup" overrideModeDefault="Deny" allowDefinition="MachineToApplication" />
      </sectionGroup>
      <section name="webSocket" overrideModeDefault="Deny" />
    </sectionGroup>
  </configSections>

  <configProtectedData defaultProvider="RsaProtectedConfigurationProvider">
    <providers>
      <add name="IISWASOnlyRsaProvider" type="" description="Uses RsaCryptoServiceProvider to encrypt and decrypt" keyContainerName="iisWasKey" cspProviderName="" useMachineContainer="true" useOAEP="false" />
      <add name="AesProvider" type="Microsoft.ApplicationHost.AesProtectedConfigurationProvider" description="Uses an AES session key to encrypt and decrypt" keyContainerName="iisConfigurationKey" cspProviderName="" useOAEP="false" useMachineContainer="true" sessionKey="AQIAAA5mAAAApAAAS6Z5YP0OSlsiaa0bcRsU29hZheG5neHzP77phcrjyRCVV1aHijkfYNx/9nGWIDdnoDgx/cfFb1w/lKNKHCFWQpUX0+UsJV29EQfvCAYJ5FKWvnKBXr5tOeBX6mZyemD2OOHYmjuke8Xhm/lO9jAqE1/lm4/34p+zQj3L2d6h32Y=" />
      <add name="IISWASOnlyAesProvider" type="Microsoft.ApplicationHost.AesProtectedConfigurationProvider" description="Uses an AES session key to encrypt and decrypt" keyContainerName="iisWasKey" cspProviderName="" useOAEP="false" useMachineContainer="true" sessionKey="AQIAAA5mAAAApAAA5FrKYkkxSDM1ipKUTmcr3tpTC7VM4WidghCoPUDzSYnwk+74oFjWCHl+3LpwwnoRgiEU9w5nHxy4UY6TSIoK2YrH/WYb7JQ65pZsu16Z98ISjGRSlr7rhqSc9tHOvjbGjNklNFBX9chp//ahKJaa4raYFR7DOjDc1RASwvh9HbA=" />
    </providers>
  </configProtectedData>

  <system.applicationHost>

    <applicationPools>
		<add name="AppPool{{.Port}}" managedRuntimeVersion="v4.0" managedPipelineMode="Integrated" CLRConfigFile="{{.AspnetConfig}}" autoStart="true" startMode="AlwaysRunning" />
    </applicationPools>

    <listenerAdapters>
      <add name="http" />
    </listenerAdapters>

    <sites>
      <siteDefaults>
        <logFile logFormat="W3C" directory="%SystemDrive%\inetpub\logs\LogFiles" />
        <traceFailedRequestsLogging directory="%SystemDrive%\inetpub\logs\FailedReqLogFiles" enabled="true" maxLogFileSizeKB="1024" maxLogFiles="20" />
      </siteDefaults>
      <applicationDefaults applicationPool="AppPool{{.Port}}" />
      <virtualDirectoryDefaults allowSubDirConfig="true" />
			<site name="IronFoundrySite" id="1" serverAutoStart="true">
        <application path="/" applicationPool="AppPool{{.Port}}">
          <virtualDirectory path="/" physicalPath="{{.RootPath}}" />
        </application>
        <bindings>
          <binding protocol="http" bindingInformation="*:{{.Port}}:" />
        </bindings>
      </site>

    </sites>

    <webLimits />

  </system.applicationHost>

  <system.webServer>

    <asp />

    <caching enabled="true" enableKernelCache="true">
    </caching>

    <cgi />

    <defaultDocument enabled="true">
      <files>
        <add value="Default.htm" />
        <add value="Default.asp" />
        <add value="index.htm" />
        <add value="index.html" />
        <add value="iisstart.htm" />
        <add value="default.aspx" />
      </files>
    </defaultDocument>

    <directoryBrowse enabled="false" />

    <fastCgi />

    <globalModules>
      <add name="UriCacheModule" image="%windir%\System32\inetsrv\cachuri.dll" />
      <add name="FileCacheModule" image="%windir%\System32\inetsrv\cachfile.dll" />
      <add name="TokenCacheModule" image="%windir%\System32\inetsrv\cachtokn.dll" />
      <add name="HttpCacheModule" image="%windir%\System32\inetsrv\cachhttp.dll" />
      <add name="StaticCompressionModule" image="%windir%\System32\inetsrv\compstat.dll" />
      <add name="DefaultDocumentModule" image="%windir%\System32\inetsrv\defdoc.dll" />
      <add name="DirectoryListingModule" image="%windir%\System32\inetsrv\dirlist.dll" />
      <add name="ProtocolSupportModule" image="%windir%\System32\inetsrv\protsup.dll" />
      <add name="StaticFileModule" image="%windir%\System32\inetsrv\static.dll" />
      <add name="AnonymousAuthenticationModule" image="%windir%\System32\inetsrv\authanon.dll" />
      <add name="RequestFilteringModule" image="%windir%\System32\inetsrv\modrqflt.dll" />
      <add name="CustomErrorModule" image="%windir%\System32\inetsrv\custerr.dll" />
      <add name="HttpLoggingModule" image="%windir%\System32\inetsrv\loghttp.dll" />
      <add name="RequestMonitorModule" image="%windir%\System32\inetsrv\iisreqs.dll" />
      <add name="IsapiModule" image="%windir%\System32\inetsrv\isapi.dll" />
      <add name="IsapiFilterModule" image="%windir%\System32\inetsrv\filter.dll" />
      <add name="ConfigurationValidationModule" image="%windir%\System32\inetsrv\validcfg.dll" />
      <add name="ManagedEngine64" image="%windir%\Microsoft.NET\Framework64\v2.0.50727\webengine.dll" preCondition="integratedMode,runtimeVersionv2.0,bitness64" />
      <add name="ManagedEngine" image="%windir%\Microsoft.NET\Framework\v2.0.50727\webengine.dll" preCondition="integratedMode,runtimeVersionv2.0,bitness32" />
      <add name="ManagedEngineV4.0_32bit" image="%windir%\Microsoft.NET\Framework\v4.0.30319\webengine4.dll" preCondition="integratedMode,runtimeVersionv4.0,bitness32" />
      <add name="ManagedEngineV4.0_64bit" image="%windir%\Microsoft.NET\Framework64\v4.0.30319\webengine4.dll" preCondition="integratedMode,runtimeVersionv4.0,bitness64" />
      <add name="CustomLoggingModule" image="%windir%\System32\inetsrv\logcust.dll" />
      <add name="TracingModule" image="%windir%\System32\inetsrv\iisetw.dll" />
      <add name="FailedRequestsTracingModule" image="%windir%\System32\inetsrv\iisfreb.dll" />

      <add name="WebSocketModule" image="%windir%\System32\inetsrv\iiswsock.dll" />
      <add name="DynamicCompressionModule" image="%windir%\System32\inetsrv\compdyn.dll" />
      <add name="HttpRedirectionModule" image="%windir%\System32\inetsrv\redirect.dll" />
      <add name="CertificateMappingAuthenticationModule" image="%windir%\System32\inetsrv\authcert.dll" />
      <add name="UrlAuthorizationModule" image="%windir%\System32\inetsrv\urlauthz.dll" />
      <add name="WindowsAuthenticationModule" image="%windir%\System32\inetsrv\authsspi.dll" />
      <add name="DigestAuthenticationModule" image="%windir%\System32\inetsrv\authmd5.dll" />
      <add name="IISCertificateMappingAuthenticationModule" image="%windir%\System32\inetsrv\authmap.dll" />
      <add name="IpRestrictionModule" image="%windir%\System32\inetsrv\iprestr.dll" />
      <add name="DynamicIpRestrictionModule" image="%windir%\System32\inetsrv\diprestr.dll" />
      <!--<add name="RequestMonitorModule" image="%windir%\System32\inetsrv\iisreqs.dll" />-->
    </globalModules>

    <httpCompression directory="%SystemDrive%\inetpub\temp\IIS Temporary Compressed Files">
      <scheme name="gzip" dll="%Windir%\system32\inetsrv\gzip.dll" />
      <staticTypes>
        <add mimeType="text/*" enabled="true" />
        <add mimeType="message/*" enabled="true" />
        <add mimeType="application/x-javascript" enabled="true" />
        <add mimeType="application/javascript" enabled="true" />
        <add mimeType="application/atom+xml" enabled="true" />
        <add mimeType="application/xaml+xml" enabled="true" />
        <add mimeType="*/*" enabled="false" />
      </staticTypes>
      <dynamicTypes>
        <add mimeType="text/*" enabled="true" />
        <add mimeType="message/*" enabled="true" />
        <add mimeType="application/x-javascript" enabled="true" />
        <add mimeType="application/javascript" enabled="true" />
        <add mimeType="*/*" enabled="false" />
      </dynamicTypes>
    </httpCompression>

    <httpErrors lockAttributes="allowAbsolutePathsWhenDelegated,defaultPath">
      <error statusCode="401" prefixLanguageFilePath="%SystemDrive%\inetpub\custerr" path="401.htm" />
      <error statusCode="403" prefixLanguageFilePath="%SystemDrive%\inetpub\custerr" path="403.htm" />
      <error statusCode="404" prefixLanguageFilePath="%SystemDrive%\inetpub\custerr" path="404.htm" />
      <error statusCode="405" prefixLanguageFilePath="%SystemDrive%\inetpub\custerr" path="405.htm" />
      <error statusCode="406" prefixLanguageFilePath="%SystemDrive%\inetpub\custerr" path="406.htm" />
      <error statusCode="412" prefixLanguageFilePath="%SystemDrive%\inetpub\custerr" path="412.htm" />
      <error statusCode="500" prefixLanguageFilePath="%SystemDrive%\inetpub\custerr" path="500.htm" />
      <error statusCode="501" prefixLanguageFilePath="%SystemDrive%\inetpub\custerr" path="501.htm" />
      <error statusCode="502" prefixLanguageFilePath="%SystemDrive%\inetpub\custerr" path="502.htm" />
    </httpErrors>

    <httpLogging dontLog="true" />

    <httpProtocol>
      <customHeaders>
        <clear />
      </customHeaders>
      <redirectHeaders>
        <clear />
      </redirectHeaders>
    </httpProtocol>

    <httpRedirect />

    <httpTracing />

    <isapiFilters>
      <filter name="ASP.Net_2.0.50727-64" path="%windir%\Microsoft.NET\Framework64\v2.0.50727\aspnet_filter.dll" enableCache="true" preCondition="bitness64,runtimeVersionv2.0" />
      <filter name="ASP.Net_2.0.50727.0" path="%windir%\Microsoft.NET\Framework\v2.0.50727\aspnet_filter.dll" enableCache="true" preCondition="bitness32,runtimeVersionv2.0" />
      <filter name="ASP.Net_2.0_for_v1.1" path="%windir%\Microsoft.NET\Framework\v2.0.50727\aspnet_filter.dll" enableCache="true" preCondition="runtimeVersionv1.1" />
      <filter name="ASP.Net_4.0_32bit" path="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_filter.dll" enableCache="true" preCondition="bitness32,runtimeVersionv4.0" />
      <filter name="ASP.Net_4.0_64bit" path="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_filter.dll" enableCache="true" preCondition="bitness64,runtimeVersionv4.0" />
    </isapiFilters>

    <odbcLogging />

    <security>

      <access sslFlags="None" />

      <applicationDependencies />

      <authentication>

        <anonymousAuthentication enabled="true" userName="IUSR" />

        <basicAuthentication />

        <clientCertificateMappingAuthentication />

        <digestAuthentication />

        <iisClientCertificateMappingAuthentication />

        <windowsAuthentication enabled="true">
          <providers>
            <add value="NTLM" />
            <add value="Negotiate" />
          </providers>
        </windowsAuthentication>
      </authentication>

      <authorization>
        <add accessType="Allow" users="*" />
      </authorization>

      <ipSecurity />

      <isapiCgiRestriction>
        <add path="%windir%\system32\inetsrv\asp.dll" allowed="true" groupId="ASP" description="Active Server Pages" />
        <add path="%windir%\Microsoft.NET\Framework64\v2.0.50727\aspnet_isapi.dll" allowed="true" groupId="ASP.NET v2.0.50727" description="ASP.NET v2.0.50727" />
        <add path="%windir%\Microsoft.NET\Framework\v2.0.50727\aspnet_isapi.dll" allowed="true" groupId="ASP.NET v2.0.50727" description="ASP.NET v2.0.50727" />
        <add path="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" allowed="false" groupId="ASP.NET v4.0.30319" description="ASP.NET v4.0.30319" />
        <add path="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" allowed="true" groupId="ASP.NET v4.0.30319" description="ASP.NET v4.0.30319" />
      </isapiCgiRestriction>

      <requestFiltering allowDoubleEscaping='false' allowHighBitCharacters='false'>
        <denyUrlSequences>
          <add sequence='..' />
          <add sequence='./' />
          <add sequence='\' />
          <add sequence=':' />
          <add sequence='%' />
          <add sequence='&amp;' />
        </denyUrlSequences>
        <fileExtensions allowUnlisted="true" applyToWebDAV="true">
          <add fileExtension=".asa" allowed="false" />
          <add fileExtension=".asax" allowed="false" />
          <add fileExtension=".ascx" allowed="false" />
          <add fileExtension=".master" allowed="false" />
          <add fileExtension=".skin" allowed="false" />
          <add fileExtension=".browser" allowed="false" />
          <add fileExtension=".sitemap" allowed="false" />
          <add fileExtension=".config" allowed="false" />
          <add fileExtension=".cs" allowed="false" />
          <add fileExtension=".csproj" allowed="false" />
          <add fileExtension=".vb" allowed="false" />
          <add fileExtension=".vbproj" allowed="false" />
          <add fileExtension=".webinfo" allowed="false" />
          <add fileExtension=".licx" allowed="false" />
          <add fileExtension=".resx" allowed="false" />
          <add fileExtension=".resources" allowed="false" />
          <add fileExtension=".mdb" allowed="false" />
          <add fileExtension=".vjsproj" allowed="false" />
          <add fileExtension=".java" allowed="false" />
          <add fileExtension=".jsl" allowed="false" />
          <add fileExtension=".ldb" allowed="false" />
          <add fileExtension=".dsdgm" allowed="false" />
          <add fileExtension=".ssdgm" allowed="false" />
          <add fileExtension=".lsad" allowed="false" />
          <add fileExtension=".ssmap" allowed="false" />
          <add fileExtension=".cd" allowed="false" />
          <add fileExtension=".dsprototype" allowed="false" />
          <add fileExtension=".lsaprototype" allowed="false" />
          <add fileExtension=".sdm" allowed="false" />
          <add fileExtension=".sdmDocument" allowed="false" />
          <add fileExtension=".mdf" allowed="false" />
          <add fileExtension=".ldf" allowed="false" />
          <add fileExtension=".ad" allowed="false" />
          <add fileExtension=".dd" allowed="false" />
          <add fileExtension=".ldd" allowed="false" />
          <add fileExtension=".sd" allowed="false" />
          <add fileExtension=".adprototype" allowed="false" />
          <add fileExtension=".lddprototype" allowed="false" />
          <add fileExtension=".exclude" allowed="false" />
          <add fileExtension=".refresh" allowed="false" />
          <add fileExtension=".compiled" allowed="false" />
          <add fileExtension=".msgx" allowed="false" />
          <add fileExtension=".vsdisco" allowed="false" />
          <add fileExtension=".rules" allowed="false" />
        </fileExtensions>
        <requestLimits maxAllowedContentLength='2097152' maxUrl='260' maxQueryString='2048' />
        <verbs allowUnlisted="true" applyToWebDAV="true" />
        <hiddenSegments applyToWebDAV="true">
          <add segment="web.config" />
          <add segment="bin" />
          <add segment="App_code" />
          <add segment="App_GlobalResources" />
          <add segment="App_LocalResources" />
          <add segment="App_WebReferences" />
          <add segment="App_Data" />
          <add segment="App_Browsers" />
          <add segment=".iishost" />
        </hiddenSegments>
      </requestFiltering>

    </security>

    <serverRuntime />

    <serverSideInclude />

    <staticContent lockAttributes="isDocFooterFileName">
      <mimeMap fileExtension=".323" mimeType="text/h323" />
      <mimeMap fileExtension=".3g2" mimeType="video/3gpp2" />
      <mimeMap fileExtension=".3gp2" mimeType="video/3gpp2" />
      <mimeMap fileExtension=".3gp" mimeType="video/3gpp" />
      <mimeMap fileExtension=".3gpp" mimeType="video/3gpp" />
      <mimeMap fileExtension=".aac" mimeType="audio/aac" />
      <mimeMap fileExtension=".aaf" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".aca" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".accdb" mimeType="application/msaccess" />
      <mimeMap fileExtension=".accde" mimeType="application/msaccess" />
      <mimeMap fileExtension=".accdt" mimeType="application/msaccess" />
      <mimeMap fileExtension=".acx" mimeType="application/internet-property-stream" />
      <mimeMap fileExtension=".adt" mimeType="audio/vnd.dlna.adts" />
      <mimeMap fileExtension=".adts" mimeType="audio/vnd.dlna.adts" />
      <mimeMap fileExtension=".afm" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".ai" mimeType="application/postscript" />
      <mimeMap fileExtension=".aif" mimeType="audio/x-aiff" />
      <mimeMap fileExtension=".aifc" mimeType="audio/aiff" />
      <mimeMap fileExtension=".aiff" mimeType="audio/aiff" />
      <mimeMap fileExtension=".application" mimeType="application/x-ms-application" />
      <mimeMap fileExtension=".art" mimeType="image/x-jg" />
      <mimeMap fileExtension=".asd" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".asf" mimeType="video/x-ms-asf" />
      <mimeMap fileExtension=".asi" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".asm" mimeType="text/plain" />
      <mimeMap fileExtension=".asr" mimeType="video/x-ms-asf" />
      <mimeMap fileExtension=".asx" mimeType="video/x-ms-asf" />
      <mimeMap fileExtension=".atom" mimeType="application/atom+xml" />
      <mimeMap fileExtension=".au" mimeType="audio/basic" />
      <mimeMap fileExtension=".avi" mimeType="video/x-msvideo" />
      <mimeMap fileExtension=".axs" mimeType="application/olescript" />
      <mimeMap fileExtension=".bas" mimeType="text/plain" />
      <mimeMap fileExtension=".bcpio" mimeType="application/x-bcpio" />
      <mimeMap fileExtension=".bin" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".bmp" mimeType="image/bmp" />
      <mimeMap fileExtension=".c" mimeType="text/plain" />
      <mimeMap fileExtension=".cab" mimeType="application/vnd.ms-cab-compressed" />
      <mimeMap fileExtension=".calx" mimeType="application/vnd.ms-office.calx" />
      <mimeMap fileExtension=".cat" mimeType="application/vnd.ms-pki.seccat" />
      <mimeMap fileExtension=".cdf" mimeType="application/x-cdf" />
      <mimeMap fileExtension=".chm" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".class" mimeType="application/x-java-applet" />
      <mimeMap fileExtension=".clp" mimeType="application/x-msclip" />
      <mimeMap fileExtension=".cmx" mimeType="image/x-cmx" />
      <mimeMap fileExtension=".cnf" mimeType="text/plain" />
      <mimeMap fileExtension=".cod" mimeType="image/cis-cod" />
      <mimeMap fileExtension=".cpio" mimeType="application/x-cpio" />
      <mimeMap fileExtension=".cpp" mimeType="text/plain" />
      <mimeMap fileExtension=".crd" mimeType="application/x-mscardfile" />
      <mimeMap fileExtension=".crl" mimeType="application/pkix-crl" />
      <mimeMap fileExtension=".crt" mimeType="application/x-x509-ca-cert" />
      <mimeMap fileExtension=".csh" mimeType="application/x-csh" />
      <mimeMap fileExtension=".css" mimeType="text/css" />
      <mimeMap fileExtension=".csv" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".cur" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".dcr" mimeType="application/x-director" />
      <mimeMap fileExtension=".deploy" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".der" mimeType="application/x-x509-ca-cert" />
      <mimeMap fileExtension=".dib" mimeType="image/bmp" />
      <mimeMap fileExtension=".dir" mimeType="application/x-director" />
      <mimeMap fileExtension=".disco" mimeType="text/xml" />
      <mimeMap fileExtension=".dll" mimeType="application/x-msdownload" />
      <mimeMap fileExtension=".dll.config" mimeType="text/xml" />
      <mimeMap fileExtension=".dlm" mimeType="text/dlm" />
      <mimeMap fileExtension=".doc" mimeType="application/msword" />
      <mimeMap fileExtension=".docm" mimeType="application/vnd.ms-word.document.macroEnabled.12" />
      <mimeMap fileExtension=".docx" mimeType="application/vnd.openxmlformats-officedocument.wordprocessingml.document" />
      <mimeMap fileExtension=".dot" mimeType="application/msword" />
      <mimeMap fileExtension=".dotm" mimeType="application/vnd.ms-word.template.macroEnabled.12" />
      <mimeMap fileExtension=".dotx" mimeType="application/vnd.openxmlformats-officedocument.wordprocessingml.template" />
      <mimeMap fileExtension=".dsp" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".dtd" mimeType="text/xml" />
      <mimeMap fileExtension=".dvi" mimeType="application/x-dvi" />
      <mimeMap fileExtension=".dvr-ms" mimeType="video/x-ms-dvr" />
      <mimeMap fileExtension=".dwf" mimeType="drawing/x-dwf" />
      <mimeMap fileExtension=".dwp" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".dxr" mimeType="application/x-director" />
      <mimeMap fileExtension=".eml" mimeType="message/rfc822" />
      <mimeMap fileExtension=".emz" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".eot" mimeType="application/vnd.ms-fontobject" />
      <mimeMap fileExtension=".eps" mimeType="application/postscript" />
      <mimeMap fileExtension=".etx" mimeType="text/x-setext" />
      <mimeMap fileExtension=".evy" mimeType="application/envoy" />
      <mimeMap fileExtension=".exe" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".exe.config" mimeType="text/xml" />
      <mimeMap fileExtension=".fdf" mimeType="application/vnd.fdf" />
      <mimeMap fileExtension=".fif" mimeType="application/fractals" />
      <mimeMap fileExtension=".fla" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".flr" mimeType="x-world/x-vrml" />
      <mimeMap fileExtension=".flv" mimeType="video/x-flv" />
      <mimeMap fileExtension=".gif" mimeType="image/gif" />
      <mimeMap fileExtension=".gtar" mimeType="application/x-gtar" />
      <mimeMap fileExtension=".gz" mimeType="application/x-gzip" />
      <mimeMap fileExtension=".h" mimeType="text/plain" />
      <mimeMap fileExtension=".hdf" mimeType="application/x-hdf" />
      <mimeMap fileExtension=".hdml" mimeType="text/x-hdml" />
      <mimeMap fileExtension=".hhc" mimeType="application/x-oleobject" />
      <mimeMap fileExtension=".hhk" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".hhp" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".hlp" mimeType="application/winhlp" />
      <mimeMap fileExtension=".hqx" mimeType="application/mac-binhex40" />
      <mimeMap fileExtension=".hta" mimeType="application/hta" />
      <mimeMap fileExtension=".htc" mimeType="text/x-component" />
      <mimeMap fileExtension=".htm" mimeType="text/html" />
      <mimeMap fileExtension=".html" mimeType="text/html" />
      <mimeMap fileExtension=".htt" mimeType="text/webviewhtml" />
      <mimeMap fileExtension=".hxt" mimeType="text/html" />
      <mimeMap fileExtension=".ical" mimeType="text/calendar" />
      <mimeMap fileExtension=".icalendar" mimeType="text/calendar" />
      <mimeMap fileExtension=".ico" mimeType="image/x-icon" />
      <mimeMap fileExtension=".ics" mimeType="text/calendar" />
      <mimeMap fileExtension=".ief" mimeType="image/ief" />
      <mimeMap fileExtension=".ifb" mimeType="text/calendar" />
      <mimeMap fileExtension=".iii" mimeType="application/x-iphone" />
      <mimeMap fileExtension=".inf" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".ins" mimeType="application/x-internet-signup" />
      <mimeMap fileExtension=".isp" mimeType="application/x-internet-signup" />
      <mimeMap fileExtension=".IVF" mimeType="video/x-ivf" />
      <mimeMap fileExtension=".jar" mimeType="application/java-archive" />
      <mimeMap fileExtension=".java" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".jck" mimeType="application/liquidmotion" />
      <mimeMap fileExtension=".jcz" mimeType="application/liquidmotion" />
      <mimeMap fileExtension=".jfif" mimeType="image/pjpeg" />
      <mimeMap fileExtension=".jpb" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".jpe" mimeType="image/jpeg" />
      <mimeMap fileExtension=".jpeg" mimeType="image/jpeg" />
      <mimeMap fileExtension=".jpg" mimeType="image/jpeg" />
      <mimeMap fileExtension=".js" mimeType="application/javascript" />
      <mimeMap fileExtension=".json" mimeType="application/json" />
      <mimeMap fileExtension=".jsx" mimeType="text/jscript" />
      <mimeMap fileExtension=".latex" mimeType="application/x-latex" />
      <mimeMap fileExtension=".lit" mimeType="application/x-ms-reader" />
      <mimeMap fileExtension=".lpk" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".lsf" mimeType="video/x-la-asf" />
      <mimeMap fileExtension=".lsx" mimeType="video/x-la-asf" />
      <mimeMap fileExtension=".lzh" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".m13" mimeType="application/x-msmediaview" />
      <mimeMap fileExtension=".m14" mimeType="application/x-msmediaview" />
      <mimeMap fileExtension=".m1v" mimeType="video/mpeg" />
      <mimeMap fileExtension=".m2ts" mimeType="video/vnd.dlna.mpeg-tts" />
      <mimeMap fileExtension=".m3u" mimeType="audio/x-mpegurl" />
      <mimeMap fileExtension=".m4a" mimeType="audio/mp4" />
      <mimeMap fileExtension=".m4v" mimeType="video/mp4" />
      <mimeMap fileExtension=".man" mimeType="application/x-troff-man" />
      <mimeMap fileExtension=".manifest" mimeType="application/x-ms-manifest" />
      <mimeMap fileExtension=".map" mimeType="text/plain" />
      <mimeMap fileExtension=".mdb" mimeType="application/x-msaccess" />
      <mimeMap fileExtension=".mdp" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".me" mimeType="application/x-troff-me" />
      <mimeMap fileExtension=".mht" mimeType="message/rfc822" />
      <mimeMap fileExtension=".mhtml" mimeType="message/rfc822" />
      <mimeMap fileExtension=".mid" mimeType="audio/mid" />
      <mimeMap fileExtension=".midi" mimeType="audio/mid" />
      <mimeMap fileExtension=".mix" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".mmf" mimeType="application/x-smaf" />
      <mimeMap fileExtension=".mno" mimeType="text/xml" />
      <mimeMap fileExtension=".mny" mimeType="application/x-msmoney" />
      <mimeMap fileExtension=".mov" mimeType="video/quicktime" />
      <mimeMap fileExtension=".movie" mimeType="video/x-sgi-movie" />
      <mimeMap fileExtension=".mp2" mimeType="video/mpeg" />
      <mimeMap fileExtension=".mp3" mimeType="audio/mpeg" />
      <mimeMap fileExtension=".mp4" mimeType="video/mp4" />
      <mimeMap fileExtension=".mp4v" mimeType="video/mp4" />
      <mimeMap fileExtension=".mpa" mimeType="video/mpeg" />
      <mimeMap fileExtension=".mpe" mimeType="video/mpeg" />
      <mimeMap fileExtension=".mpeg" mimeType="video/mpeg" />
      <mimeMap fileExtension=".mpg" mimeType="video/mpeg" />
      <mimeMap fileExtension=".mpp" mimeType="application/vnd.ms-project" />
      <mimeMap fileExtension=".mpv2" mimeType="video/mpeg" />
      <mimeMap fileExtension=".ms" mimeType="application/x-troff-ms" />
      <mimeMap fileExtension=".msi" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".mso" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".mvb" mimeType="application/x-msmediaview" />
      <mimeMap fileExtension=".mvc" mimeType="application/x-miva-compiled" />
      <mimeMap fileExtension=".nc" mimeType="application/x-netcdf" />
      <mimeMap fileExtension=".nsc" mimeType="video/x-ms-asf" />
      <mimeMap fileExtension=".nws" mimeType="message/rfc822" />
      <mimeMap fileExtension=".ocx" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".oda" mimeType="application/oda" />
      <mimeMap fileExtension=".odc" mimeType="text/x-ms-odc" />
      <mimeMap fileExtension=".ods" mimeType="application/oleobject" />
      <mimeMap fileExtension=".oga" mimeType="audio/ogg" />
      <mimeMap fileExtension=".ogg" mimeType="video/ogg" />
      <mimeMap fileExtension=".ogv" mimeType="video/ogg" />
      <mimeMap fileExtension=".ogx" mimeType="application/ogg" />
      <mimeMap fileExtension=".one" mimeType="application/onenote" />
      <mimeMap fileExtension=".onea" mimeType="application/onenote" />
      <mimeMap fileExtension=".onetoc" mimeType="application/onenote" />
      <mimeMap fileExtension=".onetoc2" mimeType="application/onenote" />
      <mimeMap fileExtension=".onetmp" mimeType="application/onenote" />
      <mimeMap fileExtension=".onepkg" mimeType="application/onenote" />
      <mimeMap fileExtension=".osdx" mimeType="application/opensearchdescription+xml" />
      <mimeMap fileExtension=".otf" mimeType="font/otf" />
      <mimeMap fileExtension=".p10" mimeType="application/pkcs10" />
      <mimeMap fileExtension=".p12" mimeType="application/x-pkcs12" />
      <mimeMap fileExtension=".p7b" mimeType="application/x-pkcs7-certificates" />
      <mimeMap fileExtension=".p7c" mimeType="application/pkcs7-mime" />
      <mimeMap fileExtension=".p7m" mimeType="application/pkcs7-mime" />
      <mimeMap fileExtension=".p7r" mimeType="application/x-pkcs7-certreqresp" />
      <mimeMap fileExtension=".p7s" mimeType="application/pkcs7-signature" />
      <mimeMap fileExtension=".pbm" mimeType="image/x-portable-bitmap" />
      <mimeMap fileExtension=".pcx" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".pcz" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".pdf" mimeType="application/pdf" />
      <mimeMap fileExtension=".pfb" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".pfm" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".pfx" mimeType="application/x-pkcs12" />
      <mimeMap fileExtension=".pgm" mimeType="image/x-portable-graymap" />
      <mimeMap fileExtension=".pko" mimeType="application/vnd.ms-pki.pko" />
      <mimeMap fileExtension=".pma" mimeType="application/x-perfmon" />
      <mimeMap fileExtension=".pmc" mimeType="application/x-perfmon" />
      <mimeMap fileExtension=".pml" mimeType="application/x-perfmon" />
      <mimeMap fileExtension=".pmr" mimeType="application/x-perfmon" />
      <mimeMap fileExtension=".pmw" mimeType="application/x-perfmon" />
      <mimeMap fileExtension=".png" mimeType="image/png" />
      <mimeMap fileExtension=".pnm" mimeType="image/x-portable-anymap" />
      <mimeMap fileExtension=".pnz" mimeType="image/png" />
      <mimeMap fileExtension=".pot" mimeType="application/vnd.ms-powerpoint" />
      <mimeMap fileExtension=".potm" mimeType="application/vnd.ms-powerpoint.template.macroEnabled.12" />
      <mimeMap fileExtension=".potx" mimeType="application/vnd.openxmlformats-officedocument.presentationml.template" />
      <mimeMap fileExtension=".ppam" mimeType="application/vnd.ms-powerpoint.addin.macroEnabled.12" />
      <mimeMap fileExtension=".ppm" mimeType="image/x-portable-pixmap" />
      <mimeMap fileExtension=".pps" mimeType="application/vnd.ms-powerpoint" />
      <mimeMap fileExtension=".ppsm" mimeType="application/vnd.ms-powerpoint.slideshow.macroEnabled.12" />
      <mimeMap fileExtension=".ppsx" mimeType="application/vnd.openxmlformats-officedocument.presentationml.slideshow" />
      <mimeMap fileExtension=".ppt" mimeType="application/vnd.ms-powerpoint" />
      <mimeMap fileExtension=".pptm" mimeType="application/vnd.ms-powerpoint.presentation.macroEnabled.12" />
      <mimeMap fileExtension=".pptx" mimeType="application/vnd.openxmlformats-officedocument.presentationml.presentation" />
      <mimeMap fileExtension=".prf" mimeType="application/pics-rules" />
      <mimeMap fileExtension=".prm" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".prx" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".ps" mimeType="application/postscript" />
      <mimeMap fileExtension=".psd" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".psm" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".psp" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".pub" mimeType="application/x-mspublisher" />
      <mimeMap fileExtension=".qt" mimeType="video/quicktime" />
      <mimeMap fileExtension=".qtl" mimeType="application/x-quicktimeplayer" />
      <mimeMap fileExtension=".qxd" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".ra" mimeType="audio/x-pn-realaudio" />
      <mimeMap fileExtension=".ram" mimeType="audio/x-pn-realaudio" />
      <mimeMap fileExtension=".rar" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".ras" mimeType="image/x-cmu-raster" />
      <mimeMap fileExtension=".rf" mimeType="image/vnd.rn-realflash" />
      <mimeMap fileExtension=".rgb" mimeType="image/x-rgb" />
      <mimeMap fileExtension=".rm" mimeType="application/vnd.rn-realmedia" />
      <mimeMap fileExtension=".rmi" mimeType="audio/mid" />
      <mimeMap fileExtension=".roff" mimeType="application/x-troff" />
      <mimeMap fileExtension=".rpm" mimeType="audio/x-pn-realaudio-plugin" />
      <mimeMap fileExtension=".rtf" mimeType="application/rtf" />
      <mimeMap fileExtension=".rtx" mimeType="text/richtext" />
      <mimeMap fileExtension=".scd" mimeType="application/x-msschedule" />
      <mimeMap fileExtension=".sct" mimeType="text/scriptlet" />
      <mimeMap fileExtension=".sea" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".setpay" mimeType="application/set-payment-initiation" />
      <mimeMap fileExtension=".setreg" mimeType="application/set-registration-initiation" />
      <mimeMap fileExtension=".sgml" mimeType="text/sgml" />
      <mimeMap fileExtension=".sh" mimeType="application/x-sh" />
      <mimeMap fileExtension=".shar" mimeType="application/x-shar" />
      <mimeMap fileExtension=".sit" mimeType="application/x-stuffit" />
      <mimeMap fileExtension=".sldm" mimeType="application/vnd.ms-powerpoint.slide.macroEnabled.12" />
      <mimeMap fileExtension=".sldx" mimeType="application/vnd.openxmlformats-officedocument.presentationml.slide" />
      <mimeMap fileExtension=".smd" mimeType="audio/x-smd" />
      <mimeMap fileExtension=".smi" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".smx" mimeType="audio/x-smd" />
      <mimeMap fileExtension=".smz" mimeType="audio/x-smd" />
      <mimeMap fileExtension=".snd" mimeType="audio/basic" />
      <mimeMap fileExtension=".snp" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".spc" mimeType="application/x-pkcs7-certificates" />
      <mimeMap fileExtension=".spl" mimeType="application/futuresplash" />
      <mimeMap fileExtension=".spx" mimeType="audio/ogg" />
      <mimeMap fileExtension=".src" mimeType="application/x-wais-source" />
      <mimeMap fileExtension=".ssm" mimeType="application/streamingmedia" />
      <mimeMap fileExtension=".sst" mimeType="application/vnd.ms-pki.certstore" />
      <mimeMap fileExtension=".stl" mimeType="application/vnd.ms-pki.stl" />
      <mimeMap fileExtension=".sv4cpio" mimeType="application/x-sv4cpio" />
      <mimeMap fileExtension=".sv4crc" mimeType="application/x-sv4crc" />
      <mimeMap fileExtension=".svg" mimeType="image/svg+xml" />
      <mimeMap fileExtension=".svgz" mimeType="image/svg+xml" />
      <mimeMap fileExtension=".swf" mimeType="application/x-shockwave-flash" />
      <mimeMap fileExtension=".t" mimeType="application/x-troff" />
      <mimeMap fileExtension=".tar" mimeType="application/x-tar" />
      <mimeMap fileExtension=".tcl" mimeType="application/x-tcl" />
      <mimeMap fileExtension=".tex" mimeType="application/x-tex" />
      <mimeMap fileExtension=".texi" mimeType="application/x-texinfo" />
      <mimeMap fileExtension=".texinfo" mimeType="application/x-texinfo" />
      <mimeMap fileExtension=".tgz" mimeType="application/x-compressed" />
      <mimeMap fileExtension=".thmx" mimeType="application/vnd.ms-officetheme" />
      <mimeMap fileExtension=".thn" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".tif" mimeType="image/tiff" />
      <mimeMap fileExtension=".tiff" mimeType="image/tiff" />
      <mimeMap fileExtension=".toc" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".tr" mimeType="application/x-troff" />
      <mimeMap fileExtension=".trm" mimeType="application/x-msterminal" />
      <mimeMap fileExtension=".ts" mimeType="video/vnd.dlna.mpeg-tts" />
      <mimeMap fileExtension=".tsv" mimeType="text/tab-separated-values" />
      <mimeMap fileExtension=".ttf" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".tts" mimeType="video/vnd.dlna.mpeg-tts" />
      <mimeMap fileExtension=".txt" mimeType="text/plain" />
      <mimeMap fileExtension=".u32" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".uls" mimeType="text/iuls" />
      <mimeMap fileExtension=".ustar" mimeType="application/x-ustar" />
      <mimeMap fileExtension=".vbs" mimeType="text/vbscript" />
      <mimeMap fileExtension=".vcf" mimeType="text/x-vcard" />
      <mimeMap fileExtension=".vcs" mimeType="text/plain" />
      <mimeMap fileExtension=".vdx" mimeType="application/vnd.ms-visio.viewer" />
      <mimeMap fileExtension=".vml" mimeType="text/xml" />
      <mimeMap fileExtension=".vsd" mimeType="application/vnd.visio" />
      <mimeMap fileExtension=".vss" mimeType="application/vnd.visio" />
      <mimeMap fileExtension=".vst" mimeType="application/vnd.visio" />
      <mimeMap fileExtension=".vsto" mimeType="application/x-ms-vsto" />
      <mimeMap fileExtension=".vsw" mimeType="application/vnd.visio" />
      <mimeMap fileExtension=".vsx" mimeType="application/vnd.visio" />
      <mimeMap fileExtension=".vtx" mimeType="application/vnd.visio" />
      <mimeMap fileExtension=".wav" mimeType="audio/wav" />
      <mimeMap fileExtension=".wax" mimeType="audio/x-ms-wax" />
      <mimeMap fileExtension=".wbmp" mimeType="image/vnd.wap.wbmp" />
      <mimeMap fileExtension=".wcm" mimeType="application/vnd.ms-works" />
      <mimeMap fileExtension=".wdb" mimeType="application/vnd.ms-works" />
      <mimeMap fileExtension=".webm" mimeType="video/webm" />
      <mimeMap fileExtension=".wks" mimeType="application/vnd.ms-works" />
      <mimeMap fileExtension=".wm" mimeType="video/x-ms-wm" />
      <mimeMap fileExtension=".wma" mimeType="audio/x-ms-wma" />
      <mimeMap fileExtension=".wmd" mimeType="application/x-ms-wmd" />
      <mimeMap fileExtension=".wmf" mimeType="application/x-msmetafile" />
      <mimeMap fileExtension=".wml" mimeType="text/vnd.wap.wml" />
      <mimeMap fileExtension=".wmlc" mimeType="application/vnd.wap.wmlc" />
      <mimeMap fileExtension=".wmls" mimeType="text/vnd.wap.wmlscript" />
      <mimeMap fileExtension=".wmlsc" mimeType="application/vnd.wap.wmlscriptc" />
      <mimeMap fileExtension=".wmp" mimeType="video/x-ms-wmp" />
      <mimeMap fileExtension=".wmv" mimeType="video/x-ms-wmv" />
      <mimeMap fileExtension=".wmx" mimeType="video/x-ms-wmx" />
      <mimeMap fileExtension=".wmz" mimeType="application/x-ms-wmz" />
      <mimeMap fileExtension=".woff" mimeType="font/x-woff" />
      <mimeMap fileExtension=".wps" mimeType="application/vnd.ms-works" />
      <mimeMap fileExtension=".wri" mimeType="application/x-mswrite" />
      <mimeMap fileExtension=".wrl" mimeType="x-world/x-vrml" />
      <mimeMap fileExtension=".wrz" mimeType="x-world/x-vrml" />
      <mimeMap fileExtension=".wsdl" mimeType="text/xml" />
      <mimeMap fileExtension=".wtv" mimeType="video/x-ms-wtv" />
      <mimeMap fileExtension=".wvx" mimeType="video/x-ms-wvx" />
      <mimeMap fileExtension=".x" mimeType="application/directx" />
      <mimeMap fileExtension=".xaf" mimeType="x-world/x-vrml" />
      <mimeMap fileExtension=".xaml" mimeType="application/xaml+xml" />
      <mimeMap fileExtension=".xap" mimeType="application/x-silverlight-app" />
      <mimeMap fileExtension=".xbap" mimeType="application/x-ms-xbap" />
      <mimeMap fileExtension=".xbm" mimeType="image/x-xbitmap" />
      <mimeMap fileExtension=".xdr" mimeType="text/plain" />
      <mimeMap fileExtension=".xht" mimeType="application/xhtml+xml" />
      <mimeMap fileExtension=".xhtml" mimeType="application/xhtml+xml" />
      <mimeMap fileExtension=".xla" mimeType="application/vnd.ms-excel" />
      <mimeMap fileExtension=".xlam" mimeType="application/vnd.ms-excel.addin.macroEnabled.12" />
      <mimeMap fileExtension=".xlc" mimeType="application/vnd.ms-excel" />
      <mimeMap fileExtension=".xlm" mimeType="application/vnd.ms-excel" />
      <mimeMap fileExtension=".xls" mimeType="application/vnd.ms-excel" />
      <mimeMap fileExtension=".xlsb" mimeType="application/vnd.ms-excel.sheet.binary.macroEnabled.12" />
      <mimeMap fileExtension=".xlsm" mimeType="application/vnd.ms-excel.sheet.macroEnabled.12" />
      <mimeMap fileExtension=".xlsx" mimeType="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" />
      <mimeMap fileExtension=".xlt" mimeType="application/vnd.ms-excel" />
      <mimeMap fileExtension=".xltm" mimeType="application/vnd.ms-excel.template.macroEnabled.12" />
      <mimeMap fileExtension=".xltx" mimeType="application/vnd.openxmlformats-officedocument.spreadsheetml.template" />
      <mimeMap fileExtension=".xlw" mimeType="application/vnd.ms-excel" />
      <mimeMap fileExtension=".xml" mimeType="text/xml" />
      <mimeMap fileExtension=".xof" mimeType="x-world/x-vrml" />
      <mimeMap fileExtension=".xpm" mimeType="image/x-xpixmap" />
      <mimeMap fileExtension=".xps" mimeType="application/vnd.ms-xpsdocument" />
      <mimeMap fileExtension=".xsd" mimeType="text/xml" />
      <mimeMap fileExtension=".xsf" mimeType="text/xml" />
      <mimeMap fileExtension=".xsl" mimeType="text/xml" />
      <mimeMap fileExtension=".xslt" mimeType="text/xml" />
      <mimeMap fileExtension=".xsn" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".xtp" mimeType="application/octet-stream" />
      <mimeMap fileExtension=".xwd" mimeType="image/x-xwindowdump" />
      <mimeMap fileExtension=".z" mimeType="application/x-compress" />
      <mimeMap fileExtension=".zip" mimeType="application/x-zip-compressed" />
    </staticContent>

    <tracing>

      <traceProviderDefinitions>
        <add name="ASPNET" guid="{AFF081FE-0247-4275-9C4E-021F3DC1DA35}">
          <areas>
            <add name="Infrastructure" value="1" />
            <add name="Module" value="2" />
            <add name="Page" value="4" />
            <add name="AppServices" value="8" />
          </areas>
        </add>
        <add name="WWW Server" guid="{3a2a4e84-4c21-4981-ae10-3fda0d9b0f83}">
          <areas>
            <clear />
            <add name="Authentication" value="2" />
            <add name="Security" value="4" />
            <add name="Filter" value="8" />
            <add name="StaticFile" value="16" />
            <add name="CGI" value="32" />
            <add name="Compression" value="64" />
            <add name="Cache" value="128" />
            <add name="RequestNotifications" value="256" />
            <add name="Module" value="512" />
            <add name="FastCGI" value="4096" />
          </areas>
        </add>
        <add name="ASP" guid="{06b94d9a-b15e-456e-a4ef-37c984a2cb4b}">
          <areas>
            <clear />
          </areas>
        </add>
        <add name="ISAPI Extension" guid="{a1c2040e-8840-4c31-ba11-9871031a19ea}">
          <areas>
            <clear />
          </areas>
        </add>
      </traceProviderDefinitions>

      <traceFailedRequests>
        <add path="*">
          <traceAreas>
            <add provider="ASP" verbosity="Verbose" />
            <add provider="ASPNET" areas="Infrastructure,Module,Page,AppServices" verbosity="Verbose" />
            <add provider="ISAPI Extension" verbosity="Verbose" />
            <add provider="WWW Server" areas="Authentication,Security,Filter,StaticFile,CGI,Compression,Cache,RequestNotifications,Module" verbosity="Verbose" />
          </traceAreas>
          <failureDefinitions statusCodes="200-999" />
        </add>
      </traceFailedRequests>

    </tracing>

    <urlCompression />

    <validation />

  </system.webServer>

  <system.webServer>

    <modules>
      <add name="HttpCacheModule" lockItem="true" />
      <add name="StaticCompressionModule" lockItem="true" />
      <add name="DefaultDocumentModule" lockItem="true" />
      <add name="DirectoryListingModule" lockItem="true" />
      <add name="IsapiFilterModule" lockItem="true" />
      <add name="ProtocolSupportModule" lockItem="true" />
      <add name="StaticFileModule" lockItem="true" />
      <add name="AnonymousAuthenticationModule" lockItem="true" />
      <add name="WindowsAuthenticationModule" lockItem="true" />
      <add name="RequestFilteringModule" lockItem="true" />
      <add name="CustomErrorModule" lockItem="true" />
      <add name="IsapiModule" lockItem="true" />
      <add name="HttpLoggingModule" lockItem="true" />
      <add name="ConfigurationValidationModule" lockItem="true" />
      <add name="OutputCache" type="System.Web.Caching.OutputCacheModule" preCondition="managedHandler" />
      <add name="Session" type="System.Web.SessionState.SessionStateModule" preCondition="managedHandler" />
      <add name="WindowsAuthentication" type="System.Web.Security.WindowsAuthenticationModule" preCondition="managedHandler" />
      <add name="FormsAuthentication" type="System.Web.Security.FormsAuthenticationModule" preCondition="managedHandler" />
      <add name="DefaultAuthentication" type="System.Web.Security.DefaultAuthenticationModule" preCondition="managedHandler" />
      <add name="RoleManager" type="System.Web.Security.RoleManagerModule" preCondition="managedHandler" />
      <add name="UrlAuthorization" type="System.Web.Security.UrlAuthorizationModule" preCondition="managedHandler" />
      <add name="FileAuthorization" type="System.Web.Security.FileAuthorizationModule" preCondition="managedHandler" />
      <add name="AnonymousIdentification" type="System.Web.Security.AnonymousIdentificationModule" preCondition="managedHandler" />
      <add name="Profile" type="System.Web.Profile.ProfileModule" preCondition="managedHandler" />
      <add name="UrlMappingsModule" type="System.Web.UrlMappingsModule" preCondition="managedHandler" />
      <add name="ServiceModel" type="System.ServiceModel.Activation.HttpModule, System.ServiceModel, Version=3.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" preCondition="managedHandler,runtimeVersionv2.0" />
      <add name="ServiceModel-4.0" type="System.ServiceModel.Activation.ServiceHttpModule, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" preCondition="managedHandler,runtimeVersionv4.0" />
      <add name="UrlRoutingModule-4.0" type="System.Web.Routing.UrlRoutingModule" preCondition="managedHandler,runtimeVersionv4.0" />
      <add name="ScriptModule-4.0" type="System.Web.Handlers.ScriptModule, System.Web.Extensions, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" preCondition="managedHandler,runtimeVersionv4.0" />
      <add name="CustomLoggingModule" lockItem="true" />
      <add name="FailedRequestsTracingModule" lockItem="true" />
      <add name="WebSocketModule" lockItem="true" />
      <add name="HttpRedirectionModule" lockItem="true" />
      <add name="CertificateMappingAuthenticationModule" lockItem="true" />
      <add name="UrlAuthorizationModule" lockItem="true" />
      <add name="DigestAuthenticationModule" lockItem="true" />
      <add name="IISCertificateMappingAuthenticationModule" lockItem="true" />
      <add name="IpRestrictionModule" lockItem="true" />
    </modules>

    <handlers accessPolicy="Read, Script">
      <add name="ASP Classic" path="*.asp" verb="GET,HEAD,POST" modules="IsapiModule" scriptProcessor="%windir%\system32\inetsrv\asp.dll" resourceType="File" />
      <add name="ISAPI-dll" path="*.dll" verb="*" modules="IsapiModule" resourceType="File" requireAccess="Execute" allowPathInfo="true" />
      <add name="AXD-ISAPI-4.0_32bit" path="*.axd" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="PageHandlerFactory-ISAPI-4.0_32bit" path="*.aspx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="SimpleHandlerFactory-ISAPI-4.0_32bit" path="*.ashx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="WebServiceHandlerFactory-ISAPI-4.0_32bit" path="*.asmx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="HttpRemotingHandlerFactory-rem-ISAPI-4.0_32bit" path="*.rem" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="HttpRemotingHandlerFactory-soap-ISAPI-4.0_32bit" path="*.soap" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="svc-ISAPI-4.0_32bit" path="*.svc" verb="*" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="rules-ISAPI-4.0_32bit" path="*.rules" verb="*" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="xoml-ISAPI-4.0_32bit" path="*.xoml" verb="*" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="xamlx-ISAPI-4.0_32bit" path="*.xamlx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="aspq-ISAPI-4.0_32bit" path="*.aspq" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="cshtm-ISAPI-4.0_32bit" path="*.cshtm" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="cshtml-ISAPI-4.0_32bit" path="*.cshtml" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="vbhtm-ISAPI-4.0_32bit" path="*.vbhtm" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="vbhtml-ISAPI-4.0_32bit" path="*.vbhtml" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="AXD-ISAPI-4.0_64bit" path="*.axd" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="PageHandlerFactory-ISAPI-4.0_64bit" path="*.aspx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="SimpleHandlerFactory-ISAPI-4.0_64bit" path="*.ashx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="WebServiceHandlerFactory-ISAPI-4.0_64bit" path="*.asmx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="HttpRemotingHandlerFactory-rem-ISAPI-4.0_64bit" path="*.rem" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="HttpRemotingHandlerFactory-soap-ISAPI-4.0_64bit" path="*.soap" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="svc-ISAPI-4.0_64bit" path="*.svc" verb="*" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="rules-ISAPI-4.0_64bit" path="*.rules" verb="*" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="xoml-ISAPI-4.0_64bit" path="*.xoml" verb="*" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="xamlx-ISAPI-4.0_64bit" path="*.xamlx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="aspq-ISAPI-4.0_64bit" path="*.aspq" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="cshtm-ISAPI-4.0_64bit" path="*.cshtm" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="cshtml-ISAPI-4.0_64bit" path="*.cshtml" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="vbhtm-ISAPI-4.0_64bit" path="*.vbhtm" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="vbhtml-ISAPI-4.0_64bit" path="*.vbhtml" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="TraceHandler-Integrated-4.0" path="trace.axd" verb="GET,HEAD,POST,DEBUG" type="System.Web.Handlers.TraceHandler" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="WebAdminHandler-Integrated-4.0" path="WebAdmin.axd" verb="GET,DEBUG" type="System.Web.Handlers.WebAdminHandler" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="AssemblyResourceLoader-Integrated-4.0" path="WebResource.axd" verb="GET,DEBUG" type="System.Web.Handlers.AssemblyResourceLoader" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="PageHandlerFactory-Integrated-4.0" path="*.aspx" verb="GET,HEAD,POST,DEBUG" type="System.Web.UI.PageHandlerFactory" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="SimpleHandlerFactory-Integrated-4.0" path="*.ashx" verb="GET,HEAD,POST,DEBUG" type="System.Web.UI.SimpleHandlerFactory" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="WebServiceHandlerFactory-Integrated-4.0" path="*.asmx" verb="GET,HEAD,POST,DEBUG" type="System.Web.Script.Services.ScriptHandlerFactory, System.Web.Extensions, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="HttpRemotingHandlerFactory-rem-Integrated-4.0" path="*.rem" verb="GET,HEAD,POST,DEBUG" type="System.Runtime.Remoting.Channels.Http.HttpRemotingHandlerFactory, System.Runtime.Remoting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="HttpRemotingHandlerFactory-soap-Integrated-4.0" path="*.soap" verb="GET,HEAD,POST,DEBUG" type="System.Runtime.Remoting.Channels.Http.HttpRemotingHandlerFactory, System.Runtime.Remoting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="svc-Integrated-4.0" path="*.svc" verb="*" type="System.ServiceModel.Activation.ServiceHttpHandlerFactory, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="rules-Integrated-4.0" path="*.rules" verb="*" type="System.ServiceModel.Activation.ServiceHttpHandlerFactory, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="xoml-Integrated-4.0" path="*.xoml" verb="*" type="System.ServiceModel.Activation.ServiceHttpHandlerFactory, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="xamlx-Integrated-4.0" path="*.xamlx" verb="GET,HEAD,POST,DEBUG" type="System.Xaml.Hosting.XamlHttpHandlerFactory, System.Xaml.Hosting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="aspq-Integrated-4.0" path="*.aspq" verb="GET,HEAD,POST,DEBUG" type="System.Web.HttpForbiddenHandler" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="cshtm-Integrated-4.0" path="*.cshtm" verb="GET,HEAD,POST,DEBUG" type="System.Web.HttpForbiddenHandler" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="cshtml-Integrated-4.0" path="*.cshtml" verb="GET,HEAD,POST,DEBUG" type="System.Web.HttpForbiddenHandler" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="vbhtm-Integrated-4.0" path="*.vbhtm" verb="GET,HEAD,POST,DEBUG" type="System.Web.HttpForbiddenHandler" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="vbhtml-Integrated-4.0" path="*.vbhtml" verb="GET,HEAD,POST,DEBUG" type="System.Web.HttpForbiddenHandler" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="ScriptHandlerFactoryAppServices-Integrated-4.0" path="*_AppService.axd" verb="*" type="System.Web.Script.Services.ScriptHandlerFactory, System.Web.Extensions, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31BF3856AD364E35" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="ScriptResourceIntegrated-4.0" path="ScriptResource.axd" verb="GET,HEAD" type="System.Web.Handlers.ScriptResourceHandler, System.Web.Extensions, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31BF3856AD364E35" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="TraceHandler-Integrated" path="trace.axd" verb="GET,HEAD,POST,DEBUG" type="System.Web.Handlers.TraceHandler" preCondition="integratedMode" />
      <add name="WebAdminHandler-Integrated" path="WebAdmin.axd" verb="GET,DEBUG" type="System.Web.Handlers.WebAdminHandler" preCondition="integratedMode" />
      <add name="AssemblyResourceLoader-Integrated" path="WebResource.axd" verb="GET,DEBUG" type="System.Web.Handlers.AssemblyResourceLoader" preCondition="integratedMode" />
      <add name="PageHandlerFactory-Integrated" path="*.aspx" verb="GET,HEAD,POST,DEBUG" type="System.Web.UI.PageHandlerFactory" preCondition="integratedMode" />
      <add name="SimpleHandlerFactory-Integrated" path="*.ashx" verb="GET,HEAD,POST,DEBUG" type="System.Web.UI.SimpleHandlerFactory" preCondition="integratedMode" />
      <add name="WebServiceHandlerFactory-Integrated" path="*.asmx" verb="GET,HEAD,POST,DEBUG" type="System.Web.Services.Protocols.WebServiceHandlerFactory, System.Web.Services, Version=2.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" preCondition="integratedMode,runtimeVersionv2.0" />
      <add name="HttpRemotingHandlerFactory-rem-Integrated" path="*.rem" verb="GET,HEAD,POST,DEBUG" type="System.Runtime.Remoting.Channels.Http.HttpRemotingHandlerFactory, System.Runtime.Remoting, Version=2.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" preCondition="integratedMode,runtimeVersionv2.0" />
      <add name="HttpRemotingHandlerFactory-soap-Integrated" path="*.soap" verb="GET,HEAD,POST,DEBUG" type="System.Runtime.Remoting.Channels.Http.HttpRemotingHandlerFactory, System.Runtime.Remoting, Version=2.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" preCondition="integratedMode,runtimeVersionv2.0" />
      <add name="AXD-ISAPI-2.0" path="*.axd" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v2.0.50727\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv2.0,bitness32" responseBufferLimit="0" />
      <add name="PageHandlerFactory-ISAPI-2.0" path="*.aspx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v2.0.50727\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv2.0,bitness32" responseBufferLimit="0" />
      <add name="SimpleHandlerFactory-ISAPI-2.0" path="*.ashx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v2.0.50727\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv2.0,bitness32" responseBufferLimit="0" />
      <add name="WebServiceHandlerFactory-ISAPI-2.0" path="*.asmx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v2.0.50727\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv2.0,bitness32" responseBufferLimit="0" />
      <add name="HttpRemotingHandlerFactory-rem-ISAPI-2.0" path="*.rem" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v2.0.50727\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv2.0,bitness32" responseBufferLimit="0" />
      <add name="HttpRemotingHandlerFactory-soap-ISAPI-2.0" path="*.soap" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v2.0.50727\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv2.0,bitness32" responseBufferLimit="0" />
      <add name="AXD-ISAPI-2.0-64" path="*.axd" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v2.0.50727\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv2.0,bitness64" responseBufferLimit="0" />
      <add name="PageHandlerFactory-ISAPI-2.0-64" path="*.aspx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v2.0.50727\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv2.0,bitness64" responseBufferLimit="0" />
      <add name="SimpleHandlerFactory-ISAPI-2.0-64" path="*.ashx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v2.0.50727\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv2.0,bitness64" responseBufferLimit="0" />
      <add name="WebServiceHandlerFactory-ISAPI-2.0-64" path="*.asmx" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v2.0.50727\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv2.0,bitness64" responseBufferLimit="0" />
      <add name="HttpRemotingHandlerFactory-rem-ISAPI-2.0-64" path="*.rem" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v2.0.50727\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv2.0,bitness64" responseBufferLimit="0" />
      <add name="HttpRemotingHandlerFactory-soap-ISAPI-2.0-64" path="*.soap" verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v2.0.50727\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv2.0,bitness64" responseBufferLimit="0" />
      <add name="TRACEVerbHandler" path="*" verb="TRACE" modules="ProtocolSupportModule" requireAccess="None" />
      <add name="OPTIONSVerbHandler" path="*" verb="OPTIONS" modules="ProtocolSupportModule" requireAccess="None" />
      <add name="ExtensionlessUrlHandler-ISAPI-4.0_32bit" path="*." verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness32" responseBufferLimit="0" />
      <add name="ExtensionlessUrlHandler-ISAPI-4.0_64bit" path="*." verb="GET,HEAD,POST,DEBUG" modules="IsapiModule" scriptProcessor="%windir%\Microsoft.NET\Framework64\v4.0.30319\aspnet_isapi.dll" preCondition="classicMode,runtimeVersionv4.0,bitness64" responseBufferLimit="0" />
      <add name="ExtensionlessUrlHandler-Integrated-4.0" path="*." verb="GET,HEAD,POST,DEBUG" type="System.Web.Handlers.TransferRequestHandler" preCondition="integratedMode,runtimeVersionv4.0" />
      <add name="StaticFile" path="*" verb="*" modules="StaticFileModule,DefaultDocumentModule,DirectoryListingModule" resourceType="Either" requireAccess="Read" />
    </handlers>
  </system.webServer>
</configuration>
`
const WebConfig = `<?xml version="1.0" encoding="utf-8"?>
<!-- the root web configuration file -->
<configuration>
    <!--
        Using a location directive with a missing path attribute
        scopes the configuration to the entire machine.  If used in
        conjunction with allowOverride="false", it can be used to
        prevent configuration from being altered on the machine

        Administrators that want to restrict permissions granted to
        web applications should change the default Trust level and ensure
        that overrides are not allowed
    -->
    <location allowOverride="true">
        <system.web>
            <securityPolicy>
                <trustLevel name="Full" policyFile="internal" />
                <trustLevel name="High" policyFile="web_hightrust.config" />
                <trustLevel name="Medium" policyFile="web_mediumtrust.config" />
                <trustLevel name="Low"  policyFile="web_lowtrust.config" />
                <trustLevel name="Minimal" policyFile="web_minimaltrust.config" />
            </securityPolicy>
            <trust level="Full" originUrl="" />
            <fullTrustAssemblies>
                <add
                    assemblyName="Microsoft.VisualStudio.Enterprise.AspNetHelper"
                    version="11.0.0.0"
                    publicKey="002400000480000094000000060200000024000052534131000400000100010007D1FA57C4AED9F0A32E84AA0FAEFD0DE9E8FD6AEC8F87FB03766C834C99921EB23BE79AD9D5DCC1DD9AD236132102900B723CF980957FC4E177108FC607774F29E8320E92EA05ECE4E821C0A5EFE8F1645C4C0C93C1AB99285D622CAA652C1DFAD63D745D6F2DE5F17E5EAF0FC4963D261C8A12436518206DC093344D5AD293"
                    />
                <add
                    assemblyName="Microsoft.VisualStudio.Web"
                    version="11.0.0.0"
                    publicKey="002400000480000094000000060200000024000052534131000400000100010007D1FA57C4AED9F0A32E84AA0FAEFD0DE9E8FD6AEC8F87FB03766C834C99921EB23BE79AD9D5DCC1DD9AD236132102900B723CF980957FC4E177108FC607774F29E8320E92EA05ECE4E821C0A5EFE8F1645C4C0C93C1AB99285D622CAA652C1DFAD63D745D6F2DE5F17E5EAF0FC4963D261C8A12436518206DC093344D5AD293"
                    />
                <add
                    assemblyName="Microsoft.Web.Infrastructure"
                    version="1.0.0.0"
                    publicKey="0024000004800000940000000602000000240000525341310004000001000100B5FC90E7027F67871E773A8FDE8938C81DD402BA65B9201D60593E96C492651E889CC13F1415EBB53FAC1131AE0BD333C5EE6021672D9718EA31A8AEBD0DA0072F25D87DBA6FC90FFD598ED4DA35E44C398C454307E8E33B8426143DAEC9F596836F97C8F74750E5975C64E2189F45DEF46B2A2B1247ADC3652BF5C308055DA9"
                    />
                <add
                    assemblyName="System.Data.SqlServerCe"
                    version="4.0.0.0"
                    publicKey="0024000004800000940000000602000000240000525341310004000001000100272736ad6e5f9586bac2d531eabc3acc666c2f8ec879fa94f8f7b0327d2ff2ed523448f83c3d5c5dd2dfc7bc99c5286b2c125117bf5cbe242b9d41750732b2bdffe649c6efb8e5526d526fdd130095ecdb7bf210809c6cdad8824faa9ac0310ac3cba2aa0523567b2dfa7fe250b30facbd62d4ec99b94ac47c7d3b28f1f6e4c8"
                    />
            </fullTrustAssemblies>
            <partialTrustVisibleAssemblies />
        </system.web>
    </location>

    <system.codedom>
        <compilers>
            <compiler language="c#;cs;csharp" extension=".cs" warningLevel="4" type="Microsoft.CSharp.CSharpCodeProvider, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089">
                <providerOption name="CompilerVersion" value="v4.0"/>
                <providerOption name="WarnAsError" value="false"/>
            </compiler>
            <compiler language="vb;vbs;visualbasic;vbscript" extension=".vb" warningLevel="4" type="Microsoft.VisualBasic.VBCodeProvider, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089">
                <providerOption name="CompilerVersion" value="v4.0"/>
                <providerOption name="OptionInfer" value="true"/>
                <providerOption name="WarnAsError" value="false"/>
            </compiler>
        </compilers>
    </system.codedom>

    <system.net>
        <defaultProxy>
            <proxy usesystemdefault="true" />
        </defaultProxy>
    </system.net>

    <system.web>
        <authorization>
            <allow users="*" />
        </authorization>

        <browserCaps userAgentCacheKeyLength="64">
            <result type="System.Web.Mobile.MobileCapabilities, System.Web.Mobile, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" />
        </browserCaps>

        <clientTarget>
            <add alias="uplevel" userAgent="Mozilla/5.0 (compatible; MSIE 6.0; Windows NT 5.1)" />
            <add alias="downlevel" userAgent="Generic Downlevel" />
        </clientTarget>

				<compilation tempDirectory="{{.TempDirectory}}">
            <assemblies>
                <add assembly="mscorlib" />
                <add assembly="Microsoft.CSharp, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" />
                <add assembly="System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" />
                <add assembly="System.Configuration, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" />
                <add assembly="System.Web, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" />
                <add assembly="System.Data, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" />
                <add assembly="System.Web.Services, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" />
                <add assembly="System.Xml, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" />
                <add assembly="System.Drawing, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" />
                <add assembly="System.EnterpriseServices, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" />
                <add assembly="System.IdentityModel, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" />
                <add assembly="System.Runtime.Serialization, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" />
                <add assembly="System.ServiceModel, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" />
                <add assembly="System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
                <add assembly="System.ServiceModel.Web, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
                <add assembly="System.Activities, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
                <add assembly="System.ServiceModel.Activities, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
                <add assembly="System.WorkflowServices, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
                <add assembly="System.Core, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" />
                <add assembly="System.Web.Extensions, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" />
                <add assembly="System.Data.DataSetExtensions, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" />
                <add assembly="System.Xml.Linq, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" />
                <add assembly="System.ComponentModel.DataAnnotations, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
                <add assembly="System.Web.DynamicData, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
                <add assembly="System.Web.ApplicationServices, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" />
                <add assembly="*" />
            </assemblies>
            <buildProviders>
                <add extension=".aspx" type="System.Web.Compilation.PageBuildProvider" />
                <add extension=".ascx" type="System.Web.Compilation.UserControlBuildProvider" />
                <add extension=".master" type="System.Web.Compilation.MasterPageBuildProvider" />
                <add extension=".asmx" type="System.Web.Compilation.WebServiceBuildProvider" />
                <add extension=".ashx" type="System.Web.Compilation.WebHandlerBuildProvider" />
                <add extension=".soap" type="System.Web.Compilation.WebServiceBuildProvider" />
                <add extension=".resx" type="System.Web.Compilation.ResXBuildProvider" />
                <add extension=".resources" type="System.Web.Compilation.ResourcesBuildProvider" />
                <add extension=".wsdl" type="System.Web.Compilation.WsdlBuildProvider" />
                <add extension=".xsd" type="System.Web.Compilation.XsdBuildProvider" />
                <add extension=".js" type="System.Web.Compilation.ForceCopyBuildProvider" />
                <add extension=".lic" type="System.Web.Compilation.IgnoreFileBuildProvider" />
                <add extension=".licx" type="System.Web.Compilation.IgnoreFileBuildProvider" />
                <add extension=".exclude" type="System.Web.Compilation.IgnoreFileBuildProvider" />
                <add extension=".refresh" type="System.Web.Compilation.IgnoreFileBuildProvider" />
                <add extension=".edmx" type="System.Data.Entity.Design.AspNet.EntityDesignerBuildProvider" />
                <add extension=".xoml" type="System.ServiceModel.Activation.WorkflowServiceBuildProvider, System.WorkflowServices, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
                <add extension=".svc" type="System.ServiceModel.Activation.ServiceBuildProvider, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" />
                <add extension=".xamlx" type="System.Xaml.Hosting.XamlBuildProvider, System.Xaml.Hosting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" />
            </buildProviders>
            <expressionBuilders>
                <add expressionPrefix="Resources" type="System.Web.Compilation.ResourceExpressionBuilder" />
                <add expressionPrefix="ConnectionStrings" type="System.Web.Compilation.ConnectionStringsExpressionBuilder" />
                <add expressionPrefix="AppSettings" type="System.Web.Compilation.AppSettingsExpressionBuilder" />
                <add expressionPrefix="RouteUrl" type="System.Web.Compilation.RouteUrlExpressionBuilder"/>
                <add expressionPrefix="RouteValue" type="System.Web.Compilation.RouteValueExpressionBuilder"/>
            </expressionBuilders>
            <folderLevelBuildProviders>
                <add name="DataServiceBuildProvider" type="System.Data.Services.BuildProvider.DataServiceBuildProvider, System.Data.Services.Design, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089"/>
            </folderLevelBuildProviders>
        </compilation>

        <healthMonitoring>
            <bufferModes>
                <add name="Critical Notification" maxBufferSize="100" maxFlushSize="20"
                    urgentFlushThreshold="1" regularFlushInterval="Infinite" urgentFlushInterval="00:01:00"
                    maxBufferThreads="1" />
                <add name="Notification" maxBufferSize="300" maxFlushSize="20"
                    urgentFlushThreshold="1" regularFlushInterval="Infinite" urgentFlushInterval="00:01:00"
                    maxBufferThreads="1" />
                <add name="Analysis" maxBufferSize="1000" maxFlushSize="100"
                    urgentFlushThreshold="100" regularFlushInterval="00:05:00"
                    urgentFlushInterval="00:01:00" maxBufferThreads="1" />
                <add name="Logging" maxBufferSize="1000" maxFlushSize="200" urgentFlushThreshold="800"
                    regularFlushInterval="00:30:00" urgentFlushInterval="00:05:00"
                    maxBufferThreads="1" />
            </bufferModes>

            <providers>
                <add name="EventLogProvider" type="System.Web.Management.EventLogWebEventProvider,System.Web,Version=4.0.0.0,Culture=neutral,PublicKeyToken=b03f5f7f11d50a3a" />
                <add connectionStringName="LocalSqlServer" maxEventDetailsLength="1073741823"
                    buffer="false" bufferMode="Notification" name="SqlWebEventProvider"
                    type="System.Web.Management.SqlWebEventProvider,System.Web,Version=4.0.0.0,Culture=neutral,PublicKeyToken=b03f5f7f11d50a3a" />
                <add name="WmiWebEventProvider" type="System.Web.Management.WmiWebEventProvider,System.Web,Version=4.0.0.0,Culture=neutral,PublicKeyToken=b03f5f7f11d50a3a" />
            </providers>

            <profiles>
                <add name="Default" minInstances="1" maxLimit="Infinite" minInterval="00:01:00"
                    custom="" />
                <add name="Critical" minInstances="1" maxLimit="Infinite" minInterval="00:00:00"
                    custom="" />
            </profiles>

            <rules>
                <add name="All Errors Default" eventName="All Errors" provider="EventLogProvider"
                    profile="Default" minInstances="1" maxLimit="Infinite" minInterval="00:01:00"
                    custom="" />
                <add name="Failure Audits Default" eventName="Failure Audits"
                    provider="EventLogProvider" profile="Default" minInstances="1"
                    maxLimit="Infinite" minInterval="00:01:00" custom="" />
            </rules>

            <eventMappings>
                <add name="All Events" type="System.Web.Management.WebBaseEvent,System.Web,Version=4.0.0.0,Culture=neutral,PublicKeyToken=b03f5f7f11d50a3a"
                    startEventCode="0" endEventCode="2147483647" />
                <add name="Heartbeats" type="System.Web.Management.WebHeartbeatEvent,System.Web,Version=4.0.0.0,Culture=neutral,PublicKeyToken=b03f5f7f11d50a3a"
                    startEventCode="0" endEventCode="2147483647" />
                <add name="Application Lifetime Events" type="System.Web.Management.WebApplicationLifetimeEvent,System.Web,Version=4.0.0.0,Culture=neutral,PublicKeyToken=b03f5f7f11d50a3a"
                    startEventCode="0" endEventCode="2147483647" />
                <add name="Request Processing Events" type="System.Web.Management.WebRequestEvent,System.Web,Version=4.0.0.0,Culture=neutral,PublicKeyToken=b03f5f7f11d50a3a"
                    startEventCode="0" endEventCode="2147483647" />
                <add name="All Errors" type="System.Web.Management.WebBaseErrorEvent,System.Web,Version=4.0.0.0,Culture=neutral,PublicKeyToken=b03f5f7f11d50a3a"
                    startEventCode="0" endEventCode="2147483647" />
                <add name="Infrastructure Errors" type="System.Web.Management.WebErrorEvent,System.Web,Version=4.0.0.0,Culture=neutral,PublicKeyToken=b03f5f7f11d50a3a"
                    startEventCode="0" endEventCode="2147483647" />
                <add name="Request Processing Errors" type="System.Web.Management.WebRequestErrorEvent,System.Web,Version=4.0.0.0,Culture=neutral,PublicKeyToken=b03f5f7f11d50a3a"
                    startEventCode="0" endEventCode="2147483647" />
                <add name="All Audits" type="System.Web.Management.WebAuditEvent,System.Web,Version=4.0.0.0,Culture=neutral,PublicKeyToken=b03f5f7f11d50a3a"
                    startEventCode="0" endEventCode="2147483647" />
                <add name="Failure Audits" type="System.Web.Management.WebFailureAuditEvent,System.Web,Version=4.0.0.0,Culture=neutral,PublicKeyToken=b03f5f7f11d50a3a"
                    startEventCode="0" endEventCode="2147483647" />
                <add name="Success Audits" type="System.Web.Management.WebSuccessAuditEvent,System.Web,Version=4.0.0.0,Culture=neutral,PublicKeyToken=b03f5f7f11d50a3a"
                    startEventCode="0" endEventCode="2147483647" />
            </eventMappings>

        </healthMonitoring>

        <httpHandlers>
            <add path="eurl.axd" verb="*" type="System.Web.HttpNotFoundHandler" validate="True" />
            <add path="trace.axd" verb="*" type="System.Web.Handlers.TraceHandler" validate="True" />
            <add path="WebResource.axd" verb="GET" type="System.Web.Handlers.AssemblyResourceLoader" validate="True" />
            <add verb="*" path="*_AppService.axd" type="System.Web.Script.Services.ScriptHandlerFactory, System.Web.Extensions, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" validate="False" />
            <add verb="GET,HEAD" path="ScriptResource.axd" type="System.Web.Handlers.ScriptResourceHandler, System.Web.Extensions, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" validate="False"/>
            <add path="*.axd" verb="*" type="System.Web.HttpNotFoundHandler" validate="True" />
            <add path="*.aspx" verb="*" type="System.Web.UI.PageHandlerFactory" validate="True" />
            <add path="*.ashx" verb="*" type="System.Web.UI.SimpleHandlerFactory" validate="True" />
            <add path="*.asmx" verb="*" type="System.Web.Script.Services.ScriptHandlerFactory, System.Web.Extensions, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" validate="False" />
            <add path="*.rem" verb="*" type="System.Runtime.Remoting.Channels.Http.HttpRemotingHandlerFactory, System.Runtime.Remoting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" validate="False" />
            <add path="*.soap" verb="*" type="System.Runtime.Remoting.Channels.Http.HttpRemotingHandlerFactory, System.Runtime.Remoting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" validate="False" />
            <add path="*.asax" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.ascx" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.master" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.skin" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.browser" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.sitemap" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.dll.config" verb="GET,HEAD" type="System.Web.StaticFileHandler" validate="True" />
            <add path="*.exe.config" verb="GET,HEAD" type="System.Web.StaticFileHandler" validate="True" />
            <add path="*.config" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.cs" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.csproj" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.vb" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.vbproj" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.webinfo" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.licx" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.resx" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.resources" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.mdb" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.vjsproj" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.java" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.jsl" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.ldb" verb="*" type="System.Web.HttpForbiddenHandler"  validate="True" />
            <add path="*.ad" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.dd" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.ldd" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.sd" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.cd" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.adprototype" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.lddprototype" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.sdm" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.sdmDocument" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.mdf" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.ldf" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.exclude" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.refresh" verb="*" type="System.Web.HttpForbiddenHandler" validate="True" />
            <add path="*.svc" verb="*" type="System.ServiceModel.Activation.HttpHandler, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" validate="False"/>
            <add path="*.rules" verb="*" type="System.Web.HttpForbiddenHandler" validate="True"/>
            <add path="*.xoml" verb="*" type="System.ServiceModel.Activation.HttpHandler, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" validate="False"/>
            <add path="*.xamlx" verb="*" type="System.Xaml.Hosting.XamlHttpHandlerFactory, System.Xaml.Hosting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" validate="False"/>
            <add path="*.aspq" verb="*" type="System.Web.HttpForbiddenHandler" validate="True"/>
            <add path="*.cshtm" verb="*" type="System.Web.HttpForbiddenHandler" validate="True"/>
            <add path="*.cshtml" verb="*" type="System.Web.HttpForbiddenHandler" validate="True"/>
            <add path="*.vbhtm" verb="*" type="System.Web.HttpForbiddenHandler" validate="True"/>
            <add path="*.vbhtml" verb="*" type="System.Web.HttpForbiddenHandler" validate="True"/>
            <add path="*" verb="GET,HEAD,POST" type="System.Web.DefaultHttpHandler" validate="True" />
            <add path="*" verb="*" type="System.Web.HttpMethodNotAllowedHandler" validate="True" />
        </httpHandlers>

        <httpModules>
            <add name="OutputCache" type="System.Web.Caching.OutputCacheModule" />
            <add name="Session" type="System.Web.SessionState.SessionStateModule" />
            <add name="WindowsAuthentication" type="System.Web.Security.WindowsAuthenticationModule" />
            <add name="FormsAuthentication" type="System.Web.Security.FormsAuthenticationModule" />
            <add name="PassportAuthentication" type="System.Web.Security.PassportAuthenticationModule" />
            <add name="RoleManager" type="System.Web.Security.RoleManagerModule" />
            <add name="UrlAuthorization" type="System.Web.Security.UrlAuthorizationModule" />
            <add name="FileAuthorization" type="System.Web.Security.FileAuthorizationModule" />
            <add name="AnonymousIdentification" type="System.Web.Security.AnonymousIdentificationModule" />
            <add name="Profile" type="System.Web.Profile.ProfileModule" />
            <add name="ErrorHandlerModule" type="System.Web.Mobile.ErrorHandlerModule, System.Web.Mobile, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" />
            <add name="ServiceModel" type="System.ServiceModel.Activation.HttpModule, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" />
            <add name="UrlRoutingModule-4.0" type="System.Web.Routing.UrlRoutingModule" />
            <add name="ScriptModule-4.0" type="System.Web.Handlers.ScriptModule, System.Web.Extensions, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
        </httpModules>

        <mobileControls sessionStateHistorySize="6" cookielessDataDictionaryType="System.Web.Mobile.CookielessData">
            <device name="XhtmlDeviceAdapters"
                predicateClass="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlPageAdapter"
                predicateMethod="DeviceQualifies"
                pageAdapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlPageAdapter">

                <control name="System.Web.UI.MobileControls.Panel"             adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlPanelAdapter" />
                <control name="System.Web.UI.MobileControls.Form"              adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlFormAdapter" />
                <control name="System.Web.UI.MobileControls.TextBox"           adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlTextBoxAdapter" />
                <control name="System.Web.UI.MobileControls.Label"             adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlLabelAdapter" />
                <control name="System.Web.UI.MobileControls.LiteralText"       adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlLiteralTextAdapter" />
                <control name="System.Web.UI.MobileControls.Link"              adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlLinkAdapter" />
                <control name="System.Web.UI.MobileControls.Command"           adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlCommandAdapter" />
                <control name="System.Web.UI.MobileControls.PhoneCall"         adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlPhoneCallAdapter" />
                <control name="System.Web.UI.MobileControls.List"              adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlListAdapter" />
                <control name="System.Web.UI.MobileControls.SelectionList"     adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlSelectionListAdapter" />
                <control name="System.Web.UI.MobileControls.ObjectList"        adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlObjectListAdapter" />
                <control name="System.Web.UI.MobileControls.Image"             adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlImageAdapter" />
                <control name="System.Web.UI.MobileControls.ValidationSummary" adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlValidationSummaryAdapter" />
                <control name="System.Web.UI.MobileControls.Calendar"          adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlCalendarAdapter" />
                <control name="System.Web.UI.MobileControls.TextView"          adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlTextViewAdapter" />
                <control name="System.Web.UI.MobileControls.MobileControl"     adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlControlAdapter" />
                <control name="System.Web.UI.MobileControls.BaseValidator"     adapter="System.Web.UI.MobileControls.Adapters.XhtmlAdapters.XhtmlValidatorAdapter" />
            </device>
            <device name="HtmlDeviceAdapters"
                predicateClass="System.Web.UI.MobileControls.Adapters.HtmlPageAdapter"
                predicateMethod="DeviceQualifies"
                pageAdapter="System.Web.UI.MobileControls.Adapters.HtmlPageAdapter">

                <control name="System.Web.UI.MobileControls.Panel"             adapter="System.Web.UI.MobileControls.Adapters.HtmlPanelAdapter" />
                <control name="System.Web.UI.MobileControls.Form"              adapter="System.Web.UI.MobileControls.Adapters.HtmlFormAdapter" />
                <control name="System.Web.UI.MobileControls.TextBox"           adapter="System.Web.UI.MobileControls.Adapters.HtmlTextBoxAdapter" />
                <control name="System.Web.UI.MobileControls.Label"             adapter="System.Web.UI.MobileControls.Adapters.HtmlLabelAdapter" />
                <control name="System.Web.UI.MobileControls.LiteralText"       adapter="System.Web.UI.MobileControls.Adapters.HtmlLiteralTextAdapter" />
                <control name="System.Web.UI.MobileControls.Link"              adapter="System.Web.UI.MobileControls.Adapters.HtmlLinkAdapter" />
                <control name="System.Web.UI.MobileControls.Command"           adapter="System.Web.UI.MobileControls.Adapters.HtmlCommandAdapter" />
                <control name="System.Web.UI.MobileControls.PhoneCall"         adapter="System.Web.UI.MobileControls.Adapters.HtmlPhoneCallAdapter" />
                <control name="System.Web.UI.MobileControls.List"              adapter="System.Web.UI.MobileControls.Adapters.HtmlListAdapter" />
                <control name="System.Web.UI.MobileControls.SelectionList"     adapter="System.Web.UI.MobileControls.Adapters.HtmlSelectionListAdapter" />
                <control name="System.Web.UI.MobileControls.ObjectList"        adapter="System.Web.UI.MobileControls.Adapters.HtmlObjectListAdapter" />
                <control name="System.Web.UI.MobileControls.Image"             adapter="System.Web.UI.MobileControls.Adapters.HtmlImageAdapter" />
                <control name="System.Web.UI.MobileControls.BaseValidator"     adapter="System.Web.UI.MobileControls.Adapters.HtmlValidatorAdapter" />
                <control name="System.Web.UI.MobileControls.ValidationSummary" adapter="System.Web.UI.MobileControls.Adapters.HtmlValidationSummaryAdapter" />
                <control name="System.Web.UI.MobileControls.Calendar"          adapter="System.Web.UI.MobileControls.Adapters.HtmlCalendarAdapter" />
                <control name="System.Web.UI.MobileControls.TextView"          adapter="System.Web.UI.MobileControls.Adapters.HtmlTextViewAdapter" />
                <control name="System.Web.UI.MobileControls.MobileControl"     adapter="System.Web.UI.MobileControls.Adapters.HtmlControlAdapter" />
            </device>
            <device name="UpWmlDeviceAdapters"
                inheritsFrom="WmlDeviceAdapters"
                predicateClass="System.Web.UI.MobileControls.Adapters.UpWmlPageAdapter"
                predicateMethod="DeviceQualifies"
                pageAdapter="System.Web.UI.MobileControls.Adapters.UpWmlPageAdapter">
            </device>
            <device name="WmlDeviceAdapters"
                predicateClass="System.Web.UI.MobileControls.Adapters.WmlPageAdapter"
                predicateMethod="DeviceQualifies"
                pageAdapter="System.Web.UI.MobileControls.Adapters.WmlPageAdapter">

                <control name="System.Web.UI.MobileControls.Panel"             adapter="System.Web.UI.MobileControls.Adapters.WmlPanelAdapter" />
                <control name="System.Web.UI.MobileControls.Form"              adapter="System.Web.UI.MobileControls.Adapters.WmlFormAdapter" />
                <control name="System.Web.UI.MobileControls.TextBox"           adapter="System.Web.UI.MobileControls.Adapters.WmlTextBoxAdapter" />
                <control name="System.Web.UI.MobileControls.Label"             adapter="System.Web.UI.MobileControls.Adapters.WmlLabelAdapter" />
                <control name="System.Web.UI.MobileControls.LiteralText"       adapter="System.Web.UI.MobileControls.Adapters.WmlLiteralTextAdapter" />
                <control name="System.Web.UI.MobileControls.Link"              adapter="System.Web.UI.MobileControls.Adapters.WmlLinkAdapter" />
                <control name="System.Web.UI.MobileControls.Command"           adapter="System.Web.UI.MobileControls.Adapters.WmlCommandAdapter" />
                <control name="System.Web.UI.MobileControls.PhoneCall"         adapter="System.Web.UI.MobileControls.Adapters.WmlPhoneCallAdapter" />
                <control name="System.Web.UI.MobileControls.List"              adapter="System.Web.UI.MobileControls.Adapters.WmlListAdapter" />
                <control name="System.Web.UI.MobileControls.SelectionList"     adapter="System.Web.UI.MobileControls.Adapters.WmlSelectionListAdapter" />
                <control name="System.Web.UI.MobileControls.ObjectList"        adapter="System.Web.UI.MobileControls.Adapters.WmlObjectListAdapter" />
                <control name="System.Web.UI.MobileControls.Image"             adapter="System.Web.UI.MobileControls.Adapters.WmlImageAdapter" />
                <control name="System.Web.UI.MobileControls.BaseValidator"     adapter="System.Web.UI.MobileControls.Adapters.WmlValidatorAdapter" />
                <control name="System.Web.UI.MobileControls.ValidationSummary" adapter="System.Web.UI.MobileControls.Adapters.WmlValidationSummaryAdapter" />
                <control name="System.Web.UI.MobileControls.Calendar"          adapter="System.Web.UI.MobileControls.Adapters.WmlCalendarAdapter" />
                <control name="System.Web.UI.MobileControls.TextView"          adapter="System.Web.UI.MobileControls.Adapters.WmlTextViewAdapter" />
                <control name="System.Web.UI.MobileControls.MobileControl"     adapter="System.Web.UI.MobileControls.Adapters.WmlControlAdapter" />
            </device>
            <device name="ChtmlDeviceAdapters"
                inheritsFrom="HtmlDeviceAdapters"
                predicateClass="System.Web.UI.MobileControls.Adapters.ChtmlPageAdapter"
                predicateMethod="DeviceQualifies"
                pageAdapter="System.Web.UI.MobileControls.Adapters.ChtmlPageAdapter">

                <control name="System.Web.UI.MobileControls.Form"              adapter="System.Web.UI.MobileControls.Adapters.ChtmlFormAdapter" />
                <control name="System.Web.UI.MobileControls.Calendar"          adapter="System.Web.UI.MobileControls.Adapters.ChtmlCalendarAdapter" />
                <control name="System.Web.UI.MobileControls.Image"             adapter="System.Web.UI.MobileControls.Adapters.ChtmlImageAdapter" />
                <control name="System.Web.UI.MobileControls.TextBox"           adapter="System.Web.UI.MobileControls.Adapters.ChtmlTextBoxAdapter" />
                <control name="System.Web.UI.MobileControls.SelectionList"     adapter="System.Web.UI.MobileControls.Adapters.ChtmlSelectionListAdapter" />
                <control name="System.Web.UI.MobileControls.Command"           adapter="System.Web.UI.MobileControls.Adapters.ChtmlCommandAdapter" />
                <control name="System.Web.UI.MobileControls.PhoneCall"         adapter="System.Web.UI.MobileControls.Adapters.ChtmlPhoneCallAdapter" />
                <control name="System.Web.UI.MobileControls.Link"              adapter="System.Web.UI.MobileControls.Adapters.ChtmlLinkAdapter" />
            </device>
        </mobileControls>

        <pages>
            <namespaces>
                <add namespace="System" />
                <add namespace="System.Collections" />
                <add namespace="System.Collections.Generic" />
                <add namespace="System.Collections.Specialized" />
                <add namespace="System.ComponentModel.DataAnnotations" />
                <add namespace="System.Configuration" />
                <add namespace="System.Linq" />
                <add namespace="System.Text" />
                <add namespace="System.Text.RegularExpressions" />
                <add namespace="System.Web" />
                <add namespace="System.Web.Caching" />
                <add namespace="System.Web.DynamicData" />
                <add namespace="System.Web.SessionState" />
                <add namespace="System.Web.Security" />
                <add namespace="System.Web.Profile" />
                <add namespace="System.Web.UI" />
                <add namespace="System.Web.UI.WebControls" />
                <add namespace="System.Web.UI.WebControls.WebParts" />
                <add namespace="System.Web.UI.HtmlControls" />
                <add namespace="System.Xml.Linq" />
            </namespaces>

            <controls>
                <add tagPrefix="asp" namespace="System.Web.UI.WebControls.WebParts" assembly="System.Web, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" />
                <add tagPrefix="asp" namespace="System.Web.UI" assembly="System.Web.Extensions, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
                <add tagPrefix="asp" namespace="System.Web.UI.WebControls" assembly="System.Web.Extensions, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
                <add tagPrefix="asp" namespace="System.Web.UI.WebControls.Expressions" assembly="System.Web.Extensions, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
                <add tagPrefix="asp" namespace="System.Web.DynamicData" assembly="System.Web.DynamicData, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
                <add tagPrefix="asp" namespace="System.Web.UI.WebControls" assembly="System.Web.Entity, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089" />
            </controls>
        </pages>

        <protocols>
            <add name="net.tcp" processHandlerType="System.ServiceModel.WasHosting.TcpProcessProtocolHandler, System.ServiceModel.WasHosting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089"
                appDomainHandlerType="System.ServiceModel.WasHosting.TcpAppDomainProtocolHandler, System.ServiceModel.WasHosting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089"
                validate="false" />
            <add name="net.pipe" processHandlerType="System.ServiceModel.WasHosting.NamedPipeProcessProtocolHandler, System.ServiceModel.WasHosting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089"
                appDomainHandlerType="System.ServiceModel.WasHosting.NamedPipeAppDomainProtocolHandler, System.ServiceModel.WasHosting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089"
                validate="false" />
            <add name="net.msmq" processHandlerType="System.ServiceModel.WasHosting.MsmqProcessProtocolHandler, System.ServiceModel.WasHosting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089"
                appDomainHandlerType="System.ServiceModel.WasHosting.MsmqAppDomainProtocolHandler, System.ServiceModel.WasHosting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089"
                validate="false" />
            <add name="msmq.formatname" processHandlerType="System.ServiceModel.WasHosting.MsmqIntegrationProcessProtocolHandler, System.ServiceModel.WasHosting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089"
                appDomainHandlerType="System.ServiceModel.WasHosting.MsmqIntegrationAppDomainProtocolHandler, System.ServiceModel.WasHosting, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089"
                validate="false" />
        </protocols>

        <siteMap>
            <providers>
                <add siteMapFile="web.sitemap" name="AspNetXmlSiteMapProvider"
                    type="System.Web.XmlSiteMapProvider, System.Web, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" />
            </providers>
        </siteMap>

        <urlMappings enabled="true" />

        <webControls clientScriptsLocation="/aspnet_client/{0}/{1}/" />

        <webParts>
            <personalization>
                <providers>
                    <add connectionStringName="LocalSqlServer"
                        name="AspNetSqlPersonalizationProvider" type="System.Web.UI.WebControls.WebParts.SqlPersonalizationProvider, System.Web, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b03f5f7f11d50a3a" />
                </providers>

                <authorization>
                    <deny users="*" verbs="enterSharedScope" />
                    <allow users="*" verbs="modifyState" />
                </authorization>
            </personalization>

            <transformers>
                <add name="RowToFieldTransformer" type="System.Web.UI.WebControls.WebParts.RowToFieldTransformer" />
                <add name="RowToParametersTransformer" type="System.Web.UI.WebControls.WebParts.RowToParametersTransformer" />
            </transformers>
        </webParts>
    </system.web>
    <system.serviceModel>
        <serviceHostingEnvironment>
            <add name="net.tcp" transportConfigurationType="System.ServiceModel.Activation.TcpHostedTransportConfiguration, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" />
            <add name="net.pipe" transportConfigurationType="System.ServiceModel.Activation.NamedPipeHostedTransportConfiguration, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" />
            <add name="net.msmq" transportConfigurationType="System.ServiceModel.Activation.MsmqHostedTransportConfiguration, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" />
            <add name="msmq.formatname" transportConfigurationType="System.ServiceModel.Activation.MsmqIntegrationHostedTransportConfiguration, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" />
        </serviceHostingEnvironment>
    </system.serviceModel>
    <system.xaml.hosting>
        <httpHandlers>
            <add xamlRootElementType="System.ServiceModel.Activities.WorkflowService, System.ServiceModel.Activities, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" httpHandlerType="System.ServiceModel.Activities.Activation.ServiceModelActivitiesActivationHandlerAsync, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
            <add xamlRootElementType="System.Activities.Activity, System.Activities, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35" httpHandlerType="System.ServiceModel.Activities.Activation.ServiceModelActivitiesActivationHandlerAsync, System.ServiceModel.Activation, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35"/>
        </httpHandlers>
    </system.xaml.hosting>
</configuration>
`
