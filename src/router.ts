import Vue from "vue";
import VueRouter from "vue-router";
import Home from "./components/home";
import FrontPage from "./components/home/frontPage";
import TestPage from "./components/home/testPage";

Vue.use(VueRouter);

const routes = [{
	children: [{
		component: FrontPage,
		name: "home",
		path: ""
	}, {
		component: TestPage,
		name: "test",
		path: "test"
	}],
	component: Home,
	path: "/"
}];

export default new VueRouter({
	mode: "history",
	routes
});
