<template>
  <div class="side-wrapper">
      <v-btn
        class="side-button mx-2"
        fab dark color="teal"
        @click="SpinDice_2d6"
      >
        <v-icon dark>mdi-format-list-bulleted-square</v-icon>
      </v-btn>
      <v-btn 
        class="side-button mx-2"
        fab dark color="teal"
        @click="SpinDice_1d6"
      >
        <v-icon dark>mdi-format-list-bulleted-square</v-icon>
      </v-btn>
      {{ career }}
      {{ state }}
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
    SpinDice_2d6: function() {
      this.career = Math.floor(Math.random() * 10) + 1;
      this.state = Math.floor(Math.random() * 10) + 1;
      console.log(this.career);
      console.log(this.state);
    },
    SpinDice_1d6: function() {
      this.ability = Math.floor(Math.random() * 6) + 1;
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