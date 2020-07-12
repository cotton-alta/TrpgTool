<template>
  <div class="card-wrapper">
    <v-container>
      <v-row>
        <v-col>
          <v-card
            class="card-title"
            outlined
            tile
          >
            {{ headline }}
          </v-card>
          <v-row no-gutters>
            <v-card
              class="card-content"
              outlined
              tile
            >
              <textarea v-model="data.value">
              </textarea>
            </v-card>
          </v-row>
        </v-col>
      </v-row>
    </v-container>
  </div>
  <!-- /.card-wrapper -->
</template>

<script>
export default {
  props: [
    "data",
    "headline"
  ]
}
</script>

<style lang="scss" scoped>
.card {
  &-wrapper {
    // height: 80px;
    margin: 10px;
  }
  &-title {
    line-height: 40px;
    height: 40px;
    text-align: center;
  }
  &-content {
    textarea {
      height: 100%;
      width: 100%;
      border: none;
    }
    text-align: center;
    line-height: 40px;
    height: 120px;
    width: 100%;
  }
}

v-row {
  height: 100%;
}
v-col {
  height: 100%;
}
</style>