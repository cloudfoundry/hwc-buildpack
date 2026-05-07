using System;
using System.Collections;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;
using System.Collections.Generic;
using System.IO;

public partial class _Default : System.Web.UI.Page
{
    protected void Page_Load(object sender, EventArgs e)
    {
        string path = HttpContext.Current.Request.Headers["X-Original-URL"];
        switch(path) {
            case "/v1/deployment/installer/agent/windows/paas/latest?bitness=64&include=dotnet&include=process":
            case "/v1/deployment/installer/agent/windows/paas/latest?bitness=64&include=dotnet&include=process&networkZone=testzone":
                string zipFilePath = Server.MapPath("paas.zip");
                Response.ContentType = "application/zip";
                Response.AddHeader("Content-Disposition", "attachment; filename=paas.zip");
                Response.WriteFile(zipFilePath);
            break;
            case "/no-manifest":
                string zipFilePathNoManifest = Server.MapPath("paas-no-manifest.zip");
                Response.ContentType = "application/zip";
                Response.AddHeader("Content-Disposition", "attachment; filename=paas.zip");
                Response.WriteFile(zipFilePathNoManifest);
            break;
            case "/v1/deployment/installer/agent/processmoduleconfig":
                string configPath = Server.MapPath("fake_config.json");
                Response.ContentType = "application/json";
                Response.WriteFile(configPath);  
            break;
            default:
                Response.StatusCode = 404;
                break;
        }

        Response.End();
    }
}
