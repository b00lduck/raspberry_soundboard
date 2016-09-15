package templates

const MainTemplate = `
<html>
<body style="font-family: arial, helvetica">
{{range .}}
	<div style="border: 1px solid black; margin: 3px; padding: 3px; float: left">
		<div>
			<a href="/play/{{.SoundFile}}">
				{{if .HasImage}}
					<img src="images/{{.ImageFile}}">
				{{end}}
			</a>
		</div>
		<div style="padding-top: 3px;">
			played {{.Count}} times
		</div>
	</div>
{{end}}
</body>
</html>
`