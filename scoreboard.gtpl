<html>
<head>
<title></title>
</head>
<body>
<form action="/scoreboard" method="post">
    Player1:<input type="text" name="player1" value="{{.Player1}}">
    Score:<input type="text" name="score1" value="{{.Score1}}">
    <br>
    Player2:<input type="text" name="player2" value="{{.Player2}}">
    Score:<input type="text" name="score2" value="{{.Score2}}">
    <input type="submit" value="save">
</form>
</body>
</html>
