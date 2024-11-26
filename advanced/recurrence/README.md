# Recurrence Relations

```
T(n) = 2T(n/2)+Cn 
```

                                N
                            /       \
                        n/2         n/2
                        /               \
                    n/4                 n/4


```
T(n) = 3T(Floor(n/4)) + Cn^2
T(n/4) = 3T(n/16) + C(n^2/16)
```

                                N
                            /   |    \
                        n/4     n/4     n/4
                        /|\     /|\      /|\       
                    n/16 n/16