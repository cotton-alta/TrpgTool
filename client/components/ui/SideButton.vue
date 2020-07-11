<template>
  <div class="side-wrapper">
      <v-btn class="side-button mx-2" fab dark color="teal">
        <v-icon dark>mdi-format-list-bulleted-square</v-icon>
      </v-btn>
      <v-btn class="side-button mx-2" fab dark color="teal">
        <v-icon dark>mdi-format-list-bulleted-square</v-icon>
      </v-btn>
  </div>
  <!-- /.side-wrapper -->
</template>

<script>
export default {
  
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