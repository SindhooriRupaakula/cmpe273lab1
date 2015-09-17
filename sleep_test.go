package sleep
import ("time" ; "testing")

func Test_sleep(t *testing.T) {
 tBefore := time.Now()
 tSleep := 4
 mySleep(tSleep)
 tAfter := time.Now()

 if tAfter == tBefore {
   t.Error("Sleep of ", tSleep, "seconds did not work!")
 }
}
