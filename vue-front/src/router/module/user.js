const userRoutes = [
  {
    path: '/register',
    name: 'register',
    // eslint-disable-next-line import/no-unresolved
    componcent: () => import('@/views/register/Register.vue'),
  },
  {
    path: '/login',
    name: 'login',
    // eslint-disable-next-line import/no-unresolved
    componcent: () => import('@/views/login/Login.vue'),
  },
  {
    path: '/profile',
    name: 'profile',
    meta: {
      auth: true,
    },
    // eslint-disable-next-line import/no-unresolved
    // componcent: () => import('@/views/profile/Profile.vue'),
  },
];

export default userRoutes;
