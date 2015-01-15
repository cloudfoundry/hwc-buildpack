using System.Xml.Linq;
using System.Xml.XPath;

namespace WebAppServer
{
    public static class XmlExtensions
    {
        public static void SetValue(this XDocument root, string elementSelector, string attributeName, object attributeValue)
        {
            root.XPathSelectElement(elementSelector).SetAttributeValue(attributeName, attributeValue);
        }

        public static void AddToElement(this XDocument root, string elementSelector, object value)
        {
            root.XPathSelectElement(elementSelector).Add(value);
        }
    }
}
