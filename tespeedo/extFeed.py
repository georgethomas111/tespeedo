import feedparser
"""
     c     := appengine.NewContext (request)
     stage := new (challenge.Stage)
     stage.Id = 1
     stage.Value = "abcdefghijklmnoprstuvwxyz"
     stage.SetKey (fmt.Sprintf ("%d", stage.Id))
     challenge.PutElem (c, "StageStore", stage)
"""
import sys
feed = feedparser.parse (sys.argv [1])
index = 5
for item in feed ["items"] :
     
     if  (len (item ["summary"])) > 200 :

         continue
     print "stage = new (challenge.Stage)"
     print "stage.Id = " + str (index)
     index = index + 1
     print 'stage.Value = "' + item ["summary"] + '"'
     print "stage.AvgScore = " + str (len (item ["summary"])/2.5)
     print 'stage.SetKey (fmt.Sprintf ("%d", stage.Id))'
     print 'challenge.PutElem (c, "StageStore", stage)'
