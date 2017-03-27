import Vue from "vue";
import router from "./routes";
import VueMaterial from "vue-material";

import "material-design-icons/iconfont/material-icons.css";
import "vue-material/dist/vue-material.css";


Vue.use(VueMaterial);
Vue.filter('truncate', function(value) {
    var length = 70;

    if (value.length <= length) {
        return value;
    } else {
        return value.substring(0, length) + '...';            
    }
});

new Vue({
	router,
	el: "#main"
});
