using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using Net.Pkcs11Interop.Common;
using Net.Pkcs11Interop.HighLevelAPI;
using Net.Pkcs11Interop.HighLevelAPI.MechanismParams;

public class Crypto
{

   private String password = "userpin";

   private static Crypto instance;

   private Pkcs11 pkcs11;

   private Slot slot;

   private ObjectHandle key;

   public static Crypto getInstance() {
      if (instance == null)
      {
         instance = new Crypto();
         Console.Out.WriteLine("Instantiated Crypto");
      }
      return instance;
   }

   public Crypto()
   {

   }

   public void Init()
   {
      try
      {
         var dllFile = new FileInfo(@".\Bin\cryptoki.dll");
         pkcs11 = new Pkcs11(dllFile.FullName, AppType.MultiThreaded);
         List<Slot> slots = pkcs11.GetSlotList(SlotsType.WithTokenPresent);
         slot = slots[0];
         Session session = slot.OpenSession(SessionType.ReadWrite);
         session.Login(CKU.CKU_USER, password);
         List<ObjectAttribute> objectAttributes = new List<ObjectAttribute>();
         objectAttributes.Add(new ObjectAttribute(CKA.CKA_CLASS, CKO.CKO_SECRET_KEY));
         objectAttributes.Add(new ObjectAttribute(CKA.CKA_KEY_TYPE, CKK.CKK_AES));
         objectAttributes.Add(new ObjectAttribute(CKA.CKA_VALUE_LEN, 32)); 
         objectAttributes.Add(new ObjectAttribute(CKA.CKA_ENCRYPT, true));
         objectAttributes.Add(new ObjectAttribute(CKA.CKA_DECRYPT, true));
         objectAttributes.Add(new ObjectAttribute(CKA.CKA_TOKEN, false));

         Mechanism mechanism = new Mechanism(CKM.CKM_AES_KEY_GEN);

         key = session.GenerateKey(mechanism, objectAttributes);

      }
      catch (Exception e)
      {
         Console.Error.WriteLine(e.ToString());
      }
   }

   public byte[] Encrypt(byte[] data, byte[] iv)
   {
      Session session = slot.OpenSession(SessionType.ReadWrite);
      try
      {
         CkAesCtrParams aesCtrParams = new CkAesCtrParams(128, iv);
         Mechanism mechanism = new Mechanism(CKM.CKM_AES_CTR, aesCtrParams);
         return session.Encrypt(mechanism, key, data);
      }
      finally
      {
         session.CloseSession();
      }
   }

   public byte[] Decrypt(byte[] encrypted, byte[]iv)
   {
      Session session = slot.OpenSession(SessionType.ReadWrite);
      try
      {
         CkAesCtrParams aesCtrParams = new CkAesCtrParams(128, iv);
         Mechanism mechanism = new Mechanism(CKM.CKM_AES_CTR, aesCtrParams);
         return session.Decrypt(mechanism, key, encrypted);
      }
      finally
      {
         session.CloseSession();
      }
   }
}
