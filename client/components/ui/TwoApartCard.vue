<template>
<div class="card-wrapper">
  <v-container>
    <v-row>
      <v-col>
        <v-row no-gutters>
          <v-col v-for="value in data.value" :key="value.id">
            <v-card
              class="card-content"
              outlined
              tile
            >
              {{ value.name }}
            </v-card>
            <v-card
              class="card-content"
              outlined
              tile
            >
              <input type="text" v-model="value.value" />
            </v-card>
          </v-col>
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
    "data"
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
    input {
      height: 100%;
      width: 100%;
      border: none;
      text-align: center;
    }
    text-align: center;
    line-height: 40px;
    height: 40px;
  }
}

v-row {
  height: 100%;
}
v-col {
  height: 100%;
}
</style>