components {
  id: "Click_Controller"
  component: "/main/Scripts/Click_Controller.script"
}
components {
  id: "ProgressBar"
  component: "/main/GUI/ProgressBar.gui"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"egg\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/Components/egg_Atlas.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "textFactory"
  type: "factory"
  data: "prototype: \"/main/GameObjects/PlusOne.go\"\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.0\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      y: 2.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 143.91304\n"
  "}\n"
  ""
}
