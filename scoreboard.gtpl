<html>
<head>
<title></title>
</head>
<body>
<form action="/scoreboard" method="post">
    Player1:<select name="player1" >
    {{$name := .Player1}}
      {{range .List}}
    {{ if eq (.) ($name)}}
    <option value="{{.}}" selected="selected">{{.}}
    </option>
    {{else}}
    <option value="{{.}}">{{.}}
      {{end}}
      {{end}}
    </select>
    Score:<input type="text" name="score1" value="{{.Score1}}">
    <br>

    Player2:<select name="player2" >
    {{$name := .Player2}}
      {{range .List}}
    {{ if eq (.) ($name)}}
    <option value="{{.}}" selected="selected">{{.}}
    </option>
    {{else}}
    <option value="{{.}}">{{.}}
    {{end}}
      {{end}}
    </select>
    Score:<input type="text" name="score2" value="{{.Score2}}">
    <input type="submit" value="save">
</select>
</form>
</body>
</html>
