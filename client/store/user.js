export const strict = false;

export const state = ({
  user: null
});

export const mutations = {
  userChange(state, value) {
    state.user = { ...value };
  }
};

export const actions = {
  userChange({ commit }, value) {
    commit("userChange", value);
  }
};

export const getters = {

};