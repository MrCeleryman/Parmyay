import {Vue, Component, Watch} from "av-ts";
import {LocalisedStrings} from "../../../util/localisedStrings";

@Component({
	template: require("./frontPage.html")
})
export default class FrontPage extends Vue {
	public localisations = LocalisedStrings.currentLocalisation;
}
