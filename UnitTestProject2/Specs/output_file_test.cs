using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using NSpec;

namespace TailorTest
{
    class describe_the_contents_of_the_output_tgz : nspec
    {
        void given_true_is_true()
        {
            it["true is true"] = () => true.should_be_true();
        }

        /****
         * 
        		Describe("the contents of the output tgz", func() { 
118 			var files []string 
119 
 
120 			JustBeforeEach(func() { 
121 				result, err := exec.Command("tar", "-tzf", outputDroplet).Output() 
122 				Ω(err).ShouldNot(HaveOccurred()) 
123 
 
124 				files = strings.Split(string(result), "\n") 
125 			}) 
126 
 
127 			It("should contain an /app dir with the contents of the compilation", func() { 
128 				Ω(files).Should(ContainElement("./app/")) 
129 				Ω(files).Should(ContainElement("./app/app.sh")) 
130 				Ω(files).Should(ContainElement("./app/compiled")) 
131 			}) 
132 
 
133 			It("should contain an empty /tmp directory", func() { 
134 				Ω(files).Should(ContainElement("./tmp/")) 
135 				Ω(files).ShouldNot(ContainElement(MatchRegexp("\\./tmp/.+"))) 
136 			}) 
137 
 
138 			It("should contain an empty /logs directory", func() { 
139 				Ω(files).Should(ContainElement("./logs/")) 
140 				Ω(files).ShouldNot(ContainElement(MatchRegexp("\\./logs/.+"))) 
141 			}) 
142 
 
143 			It("should contain a staging_info.yml with the detected buildpack", func() { 
144 				stagingInfo, err := exec.Command("tar", "-xzf", outputDroplet, "-O", "./staging_info.yml").Output() 
145 				Ω(err).ShouldNot(HaveOccurred()) 
146 
 
147 				expectedYAML := `detected_buildpack: Always Matching 
148 start_command: the start command 
149 ` 
150 				Ω(string(stagingInfo)).Should(Equal(expectedYAML)) 
151 			}) 
152 		}) 
         ****/

    }
}
