components {
  id: "ScoreText"
  component: "/main/Scripts/ScoreText.script"
}
embedded_components {
  id: "scoreText"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "text: \"Label\"\n"
  "font: \"/builtins/fonts/default.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  scale {
    x: 2.0
    y: 2.0
  }
}
