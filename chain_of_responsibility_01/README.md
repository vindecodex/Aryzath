# Another version of COR

In this version we are using main/root (starting point) that we can pass some methods and run them. Instead of
calling the next method and attach next functionality so on and so forth.

### first attempt
```
step1.SetNext(step2)
step2.SetNext(step3)
```

### second atttempt
```
steps.SetNext(step1)
steps.SetNext(step2)
steps.SetNext(step3)
```
