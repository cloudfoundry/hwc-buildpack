<%@ Page Language="C#" AutoEventWireup="true" CodeFile="Default.aspx.cs" Inherits="_Default" %>

<script runat="server">

   private byte[] StringToByteArray(string hex)
   {
      int NumberChars = hex.Length;
      byte[] bytes = new byte[NumberChars / 2];
      for (int i = 0; i < NumberChars; i += 2)
         bytes[i / 2] = Convert.ToByte(hex.Substring(i, 2), 16);
      return bytes;
   }

   private string ByteArrayToString(byte[] ba)
   {
      StringBuilder hex = new StringBuilder(ba.Length * 2);
      foreach (byte b in ba)
         hex.AppendFormat("{0:x2}", b);
      return hex.ToString();
   }

   private string encrypt(string plaintext, string ivHex)
   {
      try
      {
         byte[] iv = StringToByteArray(ivHex);
         Crypto crypto = Crypto.getInstance();
         byte[] data = System.Text.Encoding.UTF8.GetBytes(plaintext);
         byte[] encrypted = crypto.Encrypt(data, iv);
         return ByteArrayToString(encrypted);
      }
      catch (Exception exception)
      {
         return exception.ToString();
      }
   }
   
   private string decrypt(string data, string ivHex)
   {
      try
      {
         byte[] iv = StringToByteArray(ivHex);
         byte[] encrypted = StringToByteArray(data);
         Crypto crypto = Crypto.getInstance();
         byte[] decrypted = crypto.Decrypt(encrypted, iv);
         return System.Text.Encoding.UTF8.GetString(decrypted);
      }
      catch (Exception exception)
      {
         return exception.ToString();
      }
   }
</script>
<%
   string command = Request.QueryString["command"];
   string iv = Request.QueryString["iv"];
   string data = Request.QueryString["data"];
   string message = "";
   if (command == "encrypt")
   {
      message = encrypt(data, iv);
   } 
   else if (command == "decrypt")
   {
      message = decrypt(data, iv);
   }
   Response.Write(message);
%>
