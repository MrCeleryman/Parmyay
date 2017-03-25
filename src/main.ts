import Vue from "vue";
import router from "./routes";
import VueMaterial from "vue-material";

import "vue-material/dist/vue-material.css";
import "material-design-lite";

Vue.use(VueMaterial)

new Vue({
	router,
	el: "#main"
});
