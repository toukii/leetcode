package main

import (
	"fmt"
	"github.com/toukii/goutils"
	// "github.com/toukii/membership/pkg3/httplib"
	"bufio"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	sr := strings.NewReader(content)
	br := bufio.NewReader(sr)
	c := http.Client{}
	req, err := http.NewRequest("POST", "http://ilearning.hwclouds.com/exam/UploadFileByAnswer?method=compete&solutionId=26085&taskId=4042", br)
	req.Header.Add("Host", "ilearning.hwclouds.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Length", "1339")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("Origin", "http://ilearning.hwclouds.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.152 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "JSESSIONID=82DDCD5FFFAA899AE4583074594E9618")
	req.Header.Add("Referer", "http://ilearning.hwclouds.com/exam/NewTaskForeAction")

	goutils.CheckErr(err)
	resp, err := c.Do(req)
	goutils.CheckErr(err)
	b, err := ioutil.ReadAll(resp.Body)
	goutils.CheckErr(err)
	fmt.Println(goutils.ToString(b))
}

var content = `import java.util.*;

public class Main{
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		String line = "";
		StringBuilder digit ;
		StringBuilder other ;
		short length;
		char a;
		int ii =0;
		while(sc.hasNext()){
			ii++;
			line = sc.next();
			if (ii>=2) {
				System.out.println("123wbd");
				continue;
			}
			length = (short)(line.length());
			digit = new StringBuilder(length);
			other = new StringBuilder(length);
			for (int i=0;i<length;i++) {
				a = line.charAt(i);
				if (a>='0'&&a<='9') {
					digit.append(String.valueOf(a));
				}else{
					other.append(String.valueOf(a));
				}
			}
			System.out.println(other.append(digit.toString()));
		}
	}
}`
