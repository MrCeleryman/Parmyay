import {Vue, Component, Prop} from "av-ts";

@Component({
	template: require("./testPage.html")
})
export default class TestPage extends Vue {
	@Prop
	public LANG;
}
