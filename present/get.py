# Python
import urllib2
resp = urllib2.urlopen("https://example.com/")
print(resp.read())
resp.close()
