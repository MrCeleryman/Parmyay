<template>
	<div class="doughnut" @click="changeOption">
		<div class="switch" v-bind:class="optionClass"></div>
		<div ref="options">
		<p v-for="(option, index) in localisations" class="choice" v-bind:class="'choice' + (index)">{{ option}} </p> 
		</div>
	</div>
</template>
<style scoped>
    .doughnut {
		cursor: pointer;
		border: 7px solid rgba(66, 66, 66, 0.258824);
		border-radius: 100px;
		height: 25px;
		width: 25px;
	}

	.switch {
		position: relative;
		background-color: #fafafa;
		border-radius: 100px;
		height: 22px;
		width: 22px;
		transition: all .3s ease;
		box-shadow: 0 2px 2px 0 rgba(0,0,0,.14),
		0 3px 1px -2px rgba(0,0,0,.2),
		0 1px 5px 0 rgba(0,0,0,.12);
	}

	.option0 {
		left: 15px;
		top: 10px;
	}

	.option1 {
		left: -12px;
		top: 10px;
	}

	.option2 {
		left: 1px;
		top: -11px;
	}

	.choice {
		position: relative;
		font-size: 16px;
	}

	.choice0 {
		left: 40px;
		top: 10px;
	}

	.choice1 {
		left: -50px;
		top: -30px;
	}

	.choice2 {
		left: -10px;
		top: -140px;
	}

</style>
<script lang="ts">
    import {Vue, Component, Prop} from "av-ts";
	import {AllowedLocalisations, LocalisedStrings} from "../util/localisedStrings";

    @Component
    export default class TriSwitch extends Vue {
		changeIndex = 0;
		localisations = LocalisedStrings.getLocalisations();
		optionClass: any = {
			"option0": true,
			"option1": false,
			"option2": false
		}

		changeOption (): void {
			let allOptions = (this.$refs["options"] as HTMLElement).children;
			this.optionClass["option" + this.changeIndex] = false;
			
			this.changeIndex = ((++this.changeIndex) % allOptions.length);

			this.optionClass["option" + this.changeIndex] = true;

			let newLocalisation = (allOptions[this.changeIndex].innerHTML.trim() as AllowedLocalisations);
			LocalisedStrings.setLocalisation(newLocalisation)
			this.$emit("changeOption", newLocalisation);
		}
    }
</script>