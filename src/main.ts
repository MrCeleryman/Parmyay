import Vue from "vue";
import Main from "./main";

import "./styles/main.css";
import "material-design-lite";

new Vue({
	el: "#main",
	render: h => h(Main)
});