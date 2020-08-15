import dice from "~/plugins/dice";

export const strict = false;

export const state = ({
  user: null
});

export const mutations = {
  userChange(state, value) {
    state.user = { ...value };
  },
  abilityRoll(state) {
    Object.keys(state.user.dice.ability.value).forEach(key => {
      state.user.dice.ability.value[key].value = dice("1d6");
    });
  },
  careerRoll(state) {
    Object.keys(state.user.dice.career.value).forEach(key => {
      state.user.dice.career.value[key].value = dice("1d10");
    });
  },
  stateRoll(state) {
    Object.keys(state.user.dice.state.value).forEach(key => {
      state.user.dice.state.value[key].value = dice("1d10");
    });
  }
};

export const actions = {
  userChange({ commit }, value) {
    commit("userChange", value);
  },
  abilityRoll({ commit }) {
    commit("abilityRoll");
  },
  careerRoll({ commit }) {
    commit("careerRoll");
  },
  stateRoll({ commit }) {
    commit("stateRoll");
  }
};

export const getters = {
  getDice(state) {
    const { dice } = state.user;
    return dice;
  }
};