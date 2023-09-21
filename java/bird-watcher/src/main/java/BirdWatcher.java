
class BirdWatcher {

  private final int[] birdsPerDay;

  public BirdWatcher(int[] birdsPerDay) {
    this.birdsPerDay = birdsPerDay.clone();
  }

  public int[] getLastWeek() {
    return new int[]{0, 2, 5, 3, 7, 8, 4};
  }

  public int getToday() {
    return birdsPerDay[birdsPerDay.length - 1];
  }

  public void incrementTodaysCount() {
    birdsPerDay[birdsPerDay.length - 1]++;
  }

  public boolean hasDayWithoutBirds() {
    for (int i = 0; i < birdsPerDay.length - 1; i++) {
      if (birdsPerDay[i] == 0) {
        return true;
      }
    }
    return false;
  }

  public int getCountForFirstDays(int numberOfDays) {
    int sum = 0;
    int days = numberOfDays < birdsPerDay.length
        ? numberOfDays
        : birdsPerDay.length;

    for (int i = 0; i <= days - 1; i++) {
      sum += birdsPerDay[i];
    }

    return sum;
  }

  public int getBusyDays() {
    int sum = 0;
    for (int i = 0; i < birdsPerDay.length - 1; i++) {
      sum += birdsPerDay[i] >= 5 ? 1 : 0;
    }
    return sum;
  }
}
