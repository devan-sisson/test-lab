// https://dev.to/jsmanifest/state-design-pattern-in-javascript-32o9

function createStateApi(initialState) {
  const ACTION = Symbol('_action_')

  let actions = []
  let state = { ...initialState }
  let fns = {}
  let isUpdating = false
  let subscribers = []

  const createAction = (type, options) => {
    const action = { type, ...options }
    action[ACTION] = true
    return action
  }

  const setState = (nextState) => {
    state = nextState
  }

  const o = {
    createAction(type, handler) {
      const action = createAction(type)
      if (!fns[action.type]) fns[action.type] = handler
      actions.push(action)
      return action
    },
    getState() {
      return state
    },
    send(action, getAdditionalStateProps) {
      const oldState = state

      if (isUpdating) {
        return console.log(`Subscribers cannot update the state`)
      }

      try {
        isUpdating = true
        let newState = {
          ...oldState,
          ...getAdditionalStateProps?.(oldState),
          ...fns[action.type]?.(oldState),
        }
        setState(newState)
        subscribers.forEach((fn) => fn?.(oldState, newState, action))
      } finally {
        isUpdating = false
      }
    },
    subscribe(fn) {
      subscribers.push(fn)
    },
  }

  return o
}

const stateApi = createStateApi({ counter: 0, color: 'green' })

const changeColor = stateApi.createAction('changeColor')

const increment = stateApi.createAction('increment', function handler(state) {
  return {
    ...state,
    counter: state.counter + 1,
  }
})

stateApi.subscribe((oldState, newState) => {
  if (oldState.color !== newState.color) {
    console.log(`Color changed to ${newState.counter}`)
  }
})

stateApi.subscribe((oldState, newState) => {
  console.log(`The current counter is ${newState.counter}`)
})

let intervalRef = setInterval(() => {
  stateApi.send(increment)
  const state = stateApi.getState()
  const currentColor = state.color
  if (state.counter > 8 && currentColor !== 'red') {
    stateApi.send(changeColor, (state) => ({ ...state, color: 'red' }))
  } else if (state.counter >= 5 && currentColor !== 'orange') {
    stateApi.send(changeColor, (state) => ({ ...state, color: 'orange' }))
  } else if (state.counter < 5 && currentColor !== 'green') {
    stateApi.send(changeColor, (state) => ({ ...state, color: 'green' }))
  }
}, 1000)

setTimeout(() => {
  clearInterval(intervalRef)
  console.log(`Timer has ended`)
}, 10000)
