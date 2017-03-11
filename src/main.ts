import Vue from "vue";
import App from "./app";

import "./styles/main.css";
import "material-design-lite";

new Vue({
	el: "#app",
	render: h => h(App)
});