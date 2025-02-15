import { useContext, useState, useEffect } from 'react';
import { AppContext } from 'context/AppContext';
import { getAvailable } from 'api/getAvailable';
import { useNavigate } from "react-router-dom";

export function useLogin() {

  const [state, setState] = useState({
    username: '',
    password: '',
    available: false,
    disabled: true,
    busy: false,
  });

  const navigate = useNavigate();
  const app = useContext(AppContext);

  const updateState = (value) => {
    setState((s) => ({ ...s, ...value }));
  }

  const actions = {
    setUsername: (username) => {
      updateState({ username });
    },
    setPassword: (password) => {
      updateState({ password });
    },
    isDisabled: () => {
      if (state.username === '' || state.password === '') {
        return true
      }
      return false
    },
    onSettings: () => {
      navigate('/admin');
    },
    onLogin: async () => {
      if (!state.busy && state.username !== '' && state.password !== '') {
        updateState({ busy: true })
        try {
          await app.actions.login(state.username, state.password)
        }
        catch (err) {
          console.log(err);
          updateState({ busy: false })
          throw new Error('login failed: check your username and password');
        }
        updateState({ busy: false })
      }
    },
    onCreate: () => {
      navigate('/create');
    },
  };

  useEffect(() => {
    const count = async () => {
      try {
        const available = await getAvailable()
        updateState({ available: available !== 0 })
      }
      catch(err) {
        console.log(err);
      }
    }
    count();
    // eslint-disable-next-line
  }, [])

  return { state, actions };
}

