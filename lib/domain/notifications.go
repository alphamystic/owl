package domain


import (
  "owl/lib/utils"
)


func DeleteNotification(newsId string) error  {
  del,err := db.Prepare("DELETE FROM `prezo`.`news` WHERE (`newsid` = ?);")
  if err != nil{
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EDN with id: %s : %s",newsId,err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while deleting news.")
  }
  defer del.Close()
  var res sql.Result
  res,err = del.Exec(newsId)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := helpers.LogErrorToFile("sql",fmt.Sprintf("EDNZRA with id: %s : %s",newsId,err))
    helpers.Logerror(e)
    return errors.New("Server encountered an error while deleting news.")
  }
  return nil
}
