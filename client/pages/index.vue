<template>
  <div>
    <v-app>
    <div class="content-wrapper">
      <v-bottom-navigation
        :value="activeBtn"
        grow
        color="teal"
      >
        <v-btn>
          <span>キャラ1</span>
          <!-- <v-icon>mdi-history</v-icon> -->
        </v-btn>

        <v-btn>
          <span>キャラ2</span>
          <!-- <v-icon>mdi-heart</v-icon> -->
        </v-btn>

        <v-btn>
          <span>キャラ3</span>
          <!-- <v-icon>mdi-map-marker</v-icon> -->
        </v-btn>
      </v-bottom-navigation>
      <v-expansion-panels
        accordion
        class="panel-wrapper"
      >
        <v-expansion-panel>
          <v-expansion-panel-header>ダイス</v-expansion-panel-header>
          <v-expansion-panel-content>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
      <v-expansion-panels
        accordion
        class="panel-wrapper"
      >
        <v-expansion-panel>
          <v-expansion-panel-header>基本能力</v-expansion-panel-header>
          <v-expansion-panel-content>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
      <v-expansion-panels
        accordion
        class="panel-wrapper"
      >
        <v-expansion-panel>
          <v-expansion-panel-header>Item</v-expansion-panel-header>
          <v-expansion-panel-content>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </div>
    <!-- /.content-wrapper -->
    <div class="side-wrapper">

    </div>
    <!-- /.side-wrapper -->
    </v-app>
  </div>
</template>

<script>

export default {

}
</script>

<style lang="scss" scoped>
.content-wrapper {
  width: 80%;
  max-width: 1000px;
  min-height: 100vh;
  border-radius: 10px;
  margin: 0 auto;
}

.panel-wrapper {
  margin: 30px auto;
}
</style>