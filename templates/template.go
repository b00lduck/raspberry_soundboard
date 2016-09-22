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
				{{else}}
					{{.SoundFile}}
				{{end}}

			</a>
		</div>
		<div style="padding-top: 3px;">
			played {{.Count}} times
		</div>
	</div>
{{end}}
	<div style="border: 1px solid black; margin: 3px; padding: 3px; float: left">
		<div>
			<a href="/random">
				<img src="data:image/png;base64,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
ASAAAAAASUVORK5CYII=">
			</a>
		</div>
		<div style="padding-top: 3px;">
			&nbsp;
		</div>
	</div>
</body>
</html>
`