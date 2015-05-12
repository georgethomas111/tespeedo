function appendTable (data) {
  
  $('#rankTable').html("");
  
    var textToInsert = '';
    $.each(data, function(count, item) {
        textToInsert  += '<tr><td name="pieTD"><img src=' + item.Pic_Square + '/></td><td>' + item.Name + '</td><td>' + item.Score + '</td></tr>';
    });
    $('#rankTable').append(textToInsert);

}

// Make a json request to get the ranklist
function formRankList (users) {
  jsonData = JSON.stringify (users)
  console.log (jsonData);
  $("input[name=fbfrnds]").attr ("value", jsonData); 
  $("#sucForm").submit ();
}


function getFriendsFbIds () {
   console.log("In getFbFriends");
   FB.getLoginStatus(function(response) {
     if (response.status == "connected") {
       console.log (response)
       var accessToken = response.authResponse.accessToken;
     
     FB.api (
          {
            method: 'fql.query',
            query : "Select name,uid,pic_square FROM user WHERE is_app_user=1 AND uid IN(Select uid2 FROM friend WHERE uid1 = me()) OR uid = me()"
          }, 
            function(response) {
	      formRankList (response)
            }
        );
   }
 });
}
