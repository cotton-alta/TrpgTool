<template>
  <div class="side-wrapper">
      <v-btn
        class="side-button mx-2"
        fab dark color="teal"
        @click="abilityRoll"
      >
        <v-icon dark>mdi-format-list-bulleted-square</v-icon>
      </v-btn>
  </div>
  <!-- /.side-wrapper -->
</template>

<script>
export default {
  data() {
    return {
      career: 0,
      state: 0,
      ability: 0
    }
  },
  methods: {
    abilityRoll: function() {
      this.$store.dispatch("user/abilityRoll");
      this.$store.dispatch("user/careerRoll");
      this.$store.dispatch("user/stateRoll");
    },
    SpinDice_1d6: function() {
      // this.ability = Math.floor(Math.random() * 6) + 1;
      console.log(this.ability);
    }
  }
}
</script>

<style lang="scss" scoped>
.side {
  &-wrapper {
    display: inline;
    position: relative;
    top: 70px;
    left: 95%;
  }
  &-button {
    display: block;
    margin: 20px 0px;
  }
}
</style>